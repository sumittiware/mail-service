package cmd

import (
	"fmt"
	"log"
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

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf(fmt.Sprintf("Failed to start server: %v", err))
	}
}
