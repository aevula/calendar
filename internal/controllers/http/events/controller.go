package events

import (
	"net/http"

	"github.com/aevula/interview-hustlers-calendar/internal/logging"
	eventsService "github.com/aevula/interview-hustlers-calendar/internal/services/events"
)

type EventsController interface {
	CreateEvent(rw http.ResponseWriter, req *http.Request)
	UpdateEvent(rw http.ResponseWriter, req *http.Request)
	DeleteEvent(rw http.ResponseWriter, req *http.Request)
	ListEvents(rw http.ResponseWriter, req *http.Request)
}

type eventsController struct {
	service eventsService.EventsService
	logger  logging.Logger
}

func NewEventsController(service eventsService.EventsService, logger logging.Logger) EventsController {
	return &eventsController{service: service, logger: logger}
}
