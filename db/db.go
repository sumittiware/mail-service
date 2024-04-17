package db

import (
	"log"
	"os"
	"time"

	supabase "github.com/lengzuo/supa"
)

func InitDb() *supabase.Client {
	conn := connectToDB()

	if conn == nil {
		panic("Failed to connect to the DB")
	}

	return conn
}

func connectToDB() *supabase.Client {
	counts := 10

	dbHost := os.Getenv("DATABASE_HOST")
	apiKey := os.Getenv("APIKEY")

	for {
		connection, err := openDB(dbHost, apiKey)

		if err != nil {
			log.Println("DB is not yet ready")
		} else {
			log.Print("conncted to database!")
			return connection
		}

		if counts > 10 {
			return nil
		}
		counts++

		log.Print("Backing off for 1 sec")
		time.Sleep(1 * time.Second)
		continue
	}
}

func openDB(dbHost, apiKey string) (*supabase.Client, error) {
	conf := supabase.Config{
		ApiKey:     apiKey,
		ProjectRef: dbHost,
		Debug:      true,
	}
	return supabase.New(conf)
}
