package App

import (
	"github.com/redis/go-redis/v9"
)

type App struct {
	config Config
	Redis  *redis.Client
}

func CreateApp(config Config) *App {
	return &App{
		config: config,
		Redis:  config.ConnectToRedis(),
	}
}

func (app App) Run() {
	app.runServer()
}
