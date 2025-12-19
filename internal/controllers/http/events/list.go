package events

import (
	"net/http"

	"github.com/aevula/interview-hustlers-calendar/internal/apperrors"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/events/dto/requests"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/events/dto/responses"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/httputils"
)

func (c *eventsController) ListEvents(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	params, err := httputils.ParseBody(req.Body, requests.ListEventsRequest{})
	if err != nil {
		httputils.Error(rw, req, apperrors.BadRequest("BAD_REQUEST", err))
		return
	}

	cmd := params.ToCommand()

	events, err := c.service.List(ctx, cmd)
	if err != nil {
		c.logger.Error(err.Error())
		httputils.Error(rw, req, apperrors.Internal("INTERNAL", err))
		return
	}

	httputils.Render(rw, req, events, responses.ToListEventsResponse, http.StatusOK)
}
