package controllers

import (
	"github.com/aevula/interview-hustlers-calendar/internal/config"
	"github.com/aevula/interview-hustlers-calendar/internal/databases"
	"github.com/aevula/interview-hustlers-calendar/internal/logging"
	"github.com/go-chi/chi/v5"
)

type App interface {
	Config() config.Config
	Logger() *logging.Logger
	Db() databases.Db
}

func NewRouter(app App) chi.Router {
	router := chi.NewRouter()
	buildUsers(router, app)
	buildEvents(router, app)
	return router
}

func buildUsers(router chi.Router, app App) {
	router.Post("/api/v1/users", createUser(app))
}

func buildEvents(router chi.Router, app App) {
	router.Get("/api/v1/events", listEvents(app))
	router.Post("/api/v1/events", createEvent(app))
	router.Put("/api/v1/events", updateEvent(app))
	router.Delete("/api/v1/events", deleteEvent(app))
}
