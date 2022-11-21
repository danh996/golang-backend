package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth

const Secret = "test-jwt" // Replace <jwt-secret> with your secret key that is private to you.

func init() {
	tokenAuth = jwtauth.New("HS256", []byte(Secret), nil)
}

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Group(func(mux chi.Router) {
		// Seek, verify and validate JWT tokens
		mux.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		mux.Use(jwtauth.Authenticator)

		mux.Get("/schools", app.GetListSchools)
		mux.Post("/school", app.InsertSchool)
		mux.Put("/school", app.UpdateSchool)
		mux.Delete("/school", app.DeleteSchool)
	})

	mux.Group(func(mux chi.Router) {
		mux.Use(jwtauth.Verifier(tokenAuth))
		mux.Use(jwtauth.Authenticator)

		mux.Get("/users", app.GetListUsers)
	})

	mux.Group(func(r chi.Router) {
		mux.Post("/user", app.CreateUser)
		mux.Post("/login", app.Login)
	})

	return mux
}
