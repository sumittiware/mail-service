package cmd

import (
	"mail-service/config"
	"mail-service/db"
	data "mail-service/models"
	"sync"
)

func Start() {
	// connect to the database
	db := db.InitDb()

	// create channels

	// create waitgroup
	wg := sync.WaitGroup{}

	// set up the application config
	config.ApplicationConfig = &config.Config{
		DB:     db,
		Wait:   &wg,
		Models: data.New(db),
	}
	config.ApplicationConfig.Mailer = config.ApplicationConfig.CreateMail()
	// Listen to mails here!
	go config.ApplicationConfig.ListenForMail()
	// listen for shutdown signals
	go config.ApplicationConfig.ListenForShutdown()

	// Register all routes
	registerRoutes()
}
