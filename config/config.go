package config

import (
	"log"
	data "mail-service/models"
	"os"
	"os/signal"
	"sync"
	"syscall"

	supabase "github.com/lengzuo/supa"
)

var ApplicationConfig *Config

type Config struct {
	DB       *supabase.Client
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Wait     *sync.WaitGroup
	Models   data.Models
	Mailer   data.Mail
}

/*
 * ListenForShutdown listens for shutdown signals and shuts down the application
 */
func (app Config) ListenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.shutdown()
	os.Exit(0)
}

func (app *Config) shutdown() {
	// perform any cleanup tasks
	app.InfoLog.Println("would run cleanup tasks...")

	// block until waitgroup is empty
	app.Wait.Wait()
	app.Mailer.DoneChan <- true

	app.InfoLog.Println("closing channels and shutting down application...")
	close(app.Mailer.MailerChan)
	close(app.Mailer.ErrorChan)
	close(app.Mailer.DoneChan)
}

/*
 * ListenForMail listens for mail messages and sends them
 */
func (app *Config) ListenForMail() {
	for {
		select {
		case msg := <-app.Mailer.MailerChan:
			go app.Mailer.SendMail(msg, app.Mailer.ErrorChan)
		case err := <-app.Mailer.ErrorChan:
			app.ErrorLog.Println(err)
		case <-app.Mailer.DoneChan:
			return
		}
	}
}

func (app *Config) CreateMail() data.Mail {
	// create channels
	errorChan := make(chan error)
	mailerChan := make(chan data.Message, 100)
	mailerDoneChan := make(chan bool)

	m := data.Mail{
		Domain:      "localhost",
		Host:        "localhost",
		Port:        1025,
		Encryption:  "none",
		FromName:    "Info",
		FromAddress: "info@mycompany.com",
		Wait:        app.Wait,
		ErrorChan:   errorChan,
		MailerChan:  mailerChan,
		DoneChan:    mailerDoneChan,
	}

	return m
}
