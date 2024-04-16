package cmd

import (
	"mail-service/config"
	"mail-service/db"
	data "mail-service/models"
	"sync"
)

var app *config.Config

func Start() {

	// connect to the database
	db := db.InitDb()

	// create channels

	// create waitgroup
	wg := sync.WaitGroup{}

	// set up the application config
	app = &config.Config{
		DB:     db,
		Wait:   &wg,
		Models: data.New(db),
	}

	// listen for shutdown signals
	go app.ListenForShutdown()

	// Register all routes
	registerRoutes()
}
