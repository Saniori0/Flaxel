package main

import (
	"flaxel/internal/App"
	"flaxel/internal/Configurating"
	"github.com/redis/go-redis/v9"
)

func main() {
	app := App.CreateApp(Configurating.Config{
		RedisOptions: &redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		},
	})

	app.Run()
}
