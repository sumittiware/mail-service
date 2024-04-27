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
		_, err := w.Write([]byte("Welcome to Mailer!!"))
		if err != nil {
			log.Fatal(err)
			return
		}
	})

	apiRoutes := chi.NewRouter()

	// auth routes
	apiRoutes.Post("/register", SignUpHandler)
	apiRoutes.Post("/login", LoginHandler)
	apiRoutes.Post("/logout", LogoutHandler)

	// mailer routes
	apiRoutes.Get("/plans", GetPlansHandler)
	apiRoutes.Post("/subscribe/{id}", SubscribePlanHandler)

	// send mail service!
	apiRoutes.Post("/send", SendMailHandler)

	r.Mount("/api", apiRoutes)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf(fmt.Sprintf("Failed to start server: %v", err))
	}
}
