package main

import (
	"{{cookiecutter.module_name}}/config"
	"{{cookiecutter.module_name}}/router"
	"{{cookiecutter.module_name}}/server"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Msgf("failed to run server: %s", err.Error())
	}

	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	routes := router.NewRouter(cfg)
	server.Start(cfg, routes)
}
