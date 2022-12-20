package main

import (
	"log"

	"service/internal"
	"service/internal/app"
)

const configPath = "config.json"

func main() {
	config, err := internal.ParseFromFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	a := app.NewApp(config)
	a.Run()
}
