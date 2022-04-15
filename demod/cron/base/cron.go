package base

import (
	"fmt"
	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
	"github.com/robfig/cron/v3"
	"time"
)

const (
	DefaultMutexPrefix = "media-matrix/cron"
	DefaultMutexFator  = 0.05
)

type Handler struct {
	cron   string
	handle func()
	name   string
}

func newHandler(cron, name string, f func(), ) *Handler {
	return &Handler{
		cron:   cron,
		handle: f,
		name:   name,
	}
}

type RedisConfig struct {
	DNS      string
	Password string
}

type MutexConfig struct {
	RedisConfig *RedisConfig
	Prefix      string
	Factor      float64
}

type Cron struct {
	cronClient  *cron.Cron
	sync        *redsync.Redsync
	MutexConfig *MutexConfig
}

func NewCron(config *MutexConfig) *Cron {
	c := new(Cron)

	p := &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: 30 * time.Second,
		Dial: func() (redis.Conn, error) {
			if config.RedisConfig.Password == "" {
				return redis.Dial("tcp", config.RedisConfig.DNS)
			} else {
				return redis.Dial("tcp", config.RedisConfig.DNS, redis.DialPassword(config.RedisConfig.Password))
			}
		},
	}
	var pools []redsync.Pool
	pools = append(pools, p)
	c.sync = redsync.New(pools)

	if config.Prefix == "" {
		config.Prefix = DefaultMutexPrefix
	}
	if config.Factor <= 0 {
		config.Factor = DefaultMutexFator
	}
	c.MutexConfig = config
	c.cronClient = cron.New()
	return c
}

func (c *Cron) Register(task Task) {
	_, _ = c.cronClient.AddFunc(task.Rule(), wrapperHandle(c, newHandler(task.Rule(), task.Name(), task.Run)))
}

func (c *Cron) Run() {
	c.cronClient.Run()
}

func (c *Cron) lock(h *Handler) (bool, error) {
	schedule, err := cron.ParseStandard(h.cron)
	if err != nil {
		return false, err
	}
	now := time.Now()
	d := schedule.Next(now).Sub(now)
	d = d - time.Duration(float64(d)*c.MutexConfig.Factor)
	mutex := c.sync.NewMutex(fmt.Sprintf("%s/%s", c.MutexConfig.Prefix, h.name), redsync.SetExpiry(d), redsync.SetTries(1))
	if err := mutex.Lock(); err != nil {
		return false, err
	}

	//log.Printf("queue will locking still:%s %s %s\n", h.cron, h.name, schedule.Next(now))
	return true, nil
}

func wrapperHandle(c *Cron, h *Handler) func() {
	return func() {
		//log.Printf("start run queue:%s %s\n", h.cron, h.handle.Name())
		s, err := c.lock(h)
		if err != nil {
			//log.Printf("can't run queue:%s %s %s\n", h.cron, h.name, err.Error())
			return
		}
		if !s {
			//log.Printf("queue done with other processor:%s %s\n", h.cron, h.handle.Name())
			return
		}
		h.handle()
		//log.Printf("queue done:%s %s", h.cron, h.handle.Name())
	}
}
