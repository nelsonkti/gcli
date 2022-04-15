package base

var defaultCron *Cron

func StartCronTab(redisAddress string, redisPassword string) *Cron {
	return InitDefaultCron(&MutexConfig{
		RedisConfig: &RedisConfig{
			DNS:      redisAddress,
			Password: redisPassword,
		},
		Prefix: "cron",
		Factor: 0.01,
	})
}

func InitDefaultCron(config *MutexConfig) *Cron {
	if defaultCron != nil {
		panic("defaultCron init twice.")
	}
	defaultCron = NewCron(config)
	return defaultCron
}

func Register(c string, f func()) {
	if defaultCron == nil {
		panic("can't register cron before InitDefaultCron")
	}
	Register(c, f)
}

func Run() {
	if defaultCron == nil {
		panic("can't run cron before InitDefaultCron")
	}
	Run()
}
