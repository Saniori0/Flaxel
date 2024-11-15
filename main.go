package main

import (
	"Flaxel/internal/App"
	"Flaxel/internal/Configurating"
	"github.com/redis/go-redis/v9"
	"os"
	"strings"
)

func main() {
	app := App.CreateApp(Configurating.Config{
		RedisOptions: &redis.Options{
			Addr:     os.Args[1],
			Password: "",
		},
	})

	app.ExecCommand(strings.Join(os.Args[2:], " "))
}
