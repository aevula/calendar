package handlers

import (
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/events"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/users"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	BuildUsers(controller users.UsersController)
	BuildEvents(controller events.EventsController)
	Done() chi.Router
}

type router struct {
	router chi.Router
}

func NewRouter() Router {
	return &router{
		router: chi.NewRouter(),
	}
}

func (r *router) BuildUsers(controller users.UsersController) {
	r.router.Post("/api/v1/users", controller.CreateUser)
}

func (r *router) BuildEvents(controller events.EventsController) {
	r.router.Get("/api/v1/events", controller.ListEvents)
	r.router.Post("/api/v1/events", controller.CreateEvent)
	r.router.Put("/api/v1/events", controller.UpdateEvent)
	r.router.Delete("/api/v1/events", controller.DeleteEvent)
}

func (r *router) Done() chi.Router {
	return r.router
}
