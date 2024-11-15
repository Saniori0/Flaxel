package App

import (
	"github.com/redis/go-redis/v9"
)

type Config struct {
	RedisOptions *redis.Options
}

func (config Config) ConnectToRedis() *redis.Client {
	return redis.NewClient(config.RedisOptions)
}
