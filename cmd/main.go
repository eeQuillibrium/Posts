package main

import (
	"github.com/eeQuillibrium/posts/config"
	"github.com/eeQuillibrium/posts/internal/app"
	"github.com/eeQuillibrium/posts/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	log := logger.NewLogger()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("env params loading problem: %v", err)
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	appl := app.NewApp(log, cfg)

	log.Fatal(appl.Run())
}
