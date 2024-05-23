package main

import (
	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/internal/app"
	"github.com/eeQuillibrium/posts/pkg/logger"
)

func main() {
	log := logger.NewLogger()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	appl := app.NewApp(log, cfg)

	log.Fatal(appl.Run())
}
