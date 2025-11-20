package events

import (
	"net/http"

	"github.com/aevula/interview-hustlers-calendar/internal/apperrors"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/events/dto/requests"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/events/dto/responses"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/httputils"
)

func (c *eventsController) CreateEvent(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	params, err := httputils.ParseBody(req.Body, requests.CreateEventRequest{})
	if err != nil {
		httputils.Error(rw, req, apperrors.BadRequest("BAD_REQUEST", err))
		return
	}

	cmd, err := params.ToCommand()
	if err != nil {
		httputils.Error(rw, req, apperrors.BadRequest("BAD_REQUEST", err))
		return
	}

	event, err := c.service.Create(ctx, cmd)
	if err != nil {
		c.logger.Error(err.Error())
		httputils.Error(rw, req, apperrors.Internal("INTERNAL", err))
		return
	}

	c.logger.Info("Created Event", c.logger.Int("ID", int(event.ID)))
	httputils.Render(rw, req, event, responses.ToCreateEventResponse, http.StatusCreated)
}
