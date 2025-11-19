package events

import (
	"net/http"

	"github.com/aevula/interview-hustlers-calendar/internal/logging"
	"github.com/aevula/interview-hustlers-calendar/internal/services/events"
)

type EventsController interface {
	CreateEvent() http.HandlerFunc
	UpdateEvent() http.HandlerFunc
	DeleteEvent() http.HandlerFunc
	ListEvents() http.HandlerFunc
}

type eventsController struct {
	service events.EventsService
	logger  logging.Logger
}

func NewEventsController(service events.EventsService, logger logging.Logger) EventsController {
	return &eventsController{service: service, logger: logger}
}
