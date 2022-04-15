package db

import (
	"demod/lib/logger"
	"github.com/go-redis/redis"
	"sync"
)

//redis 服务
var redisDatabases sync.Map

func ConnectRedis(address string, password string, db int, name string) *redis.Client {
	redisClient := redis.NewClient(
		&redis.Options{
			Addr:     address,
			Password: password,
			DB:       db,
		},
	)
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

	redisDatabases.Store(name, redisClient)
	return redisClient
}

func Redis(name string) *redis.Client {
	value, ok := redisDatabases.Load(name)
	if ok {
		return value.(*redis.Client)
	}
	logger.Sugar.Info("failed to connect redis database:" + name)
	panic("failed to connect redis database:" + name)
}

func RedisDefault() *redis.Client {
	return Redis("default")
}

func DisconnectRedis() {
	redisDatabases.Range(func(key, value interface{}) bool {
		defer value.(*redis.Client).Close()
		return true
	})
	logger.Sugar.Info("disconnect redis")
}
