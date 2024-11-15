package App

import (
	"flaxel/internal/Configurating"
	"github.com/redis/go-redis/v9"
)

type App struct {
	config Configurating.Config
	Redis  *redis.Client
}

func CreateApp(config Configurating.Config) *App {
	return &App{
		config: config,
		Redis:  config.ConnectToRedis(),
	}
}

func (app App) Run() {
	app.runServer()
}
