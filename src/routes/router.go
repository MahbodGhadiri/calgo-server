package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	handler "calgo-server/src/handlers"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	// global middlewares
	r.Use(middleware.Logger)

	// routes
	r.Get("/", handler.WelcomeHandler())

	return r
}
