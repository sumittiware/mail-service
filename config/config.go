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
}

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

	app.InfoLog.Println("closing channels and shutting down application...")
}
