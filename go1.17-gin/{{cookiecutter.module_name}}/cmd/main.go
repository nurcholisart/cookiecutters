package main

import (
	"log"
	"github.com/joho/godotenv"
	"{{cookiecutter.module_name}}/config"
	"{{cookiecutter.module_name}}/helpers"
	"{{cookiecutter.module_name}}/server"
)

var version string

func main() {
	_ = godotenv.Load()

	conf, err := config.DefineConfig()
	if err != nil {
		log.Fatalf("Failed to run server: %s", err.Error())
	}

	conf.Version = version

	appLog := helpers.AppLog{
		Severity: "info",
		Message:  "Hello",
	}

	appLog.PrintLog()

	server.Start(conf)
}
