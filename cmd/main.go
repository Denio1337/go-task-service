package main

import (
	"app/config"
	"app/router"
	"log"
)

func main() {
	// Create application instance
	app := router.New()

	// Run application
	log.Fatal(app.Listen(config.Get(config.EnvAppAddress)))
}
