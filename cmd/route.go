package cmd

import (
	"fmt"
	"log"
	data "mail-service/models"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func registerRoutes() {
	// create router
	r := chi.NewRouter()

	// set up middleware
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	// define application routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Mailer!!"))
	})

	// auth routes
	r.Post("/register", SignUpHandler)
	r.Post("/login", LoginHandler)
	r.Post("/logout", LogoutHandler)

	// mailer routes
	r.Get("/plans", GetPlansHandler)
	// r.Post("/plan", CreatePlanHandler)
	// r.Post("/subscribe/{id}", SubscribePlanHandler)

	// send mail service!
	r.Post("/send", SendMailHandler)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		mail := data.Mail{
			Domain:      "localhost",
			Host:        "localhost",
			Port:        1025,
			Encryption:  "none",
			FromAddress: "info@mycompany.com",
			FromName:    "info",
			ErrorChan:   make(chan error),
		}
		msg := data.Message{
			To:      "tiwaresumit143@gmail.com",
			Subject: "Test Mail",
			Data:    "Hello, This is a test mail",
		}

		mail.SendMail(msg, make(chan error))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf(fmt.Sprintf("Failed to start server: %v", err))
	}
}
