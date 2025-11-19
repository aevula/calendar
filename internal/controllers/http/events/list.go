package events

import (
	"net/http"

	"github.com/aevula/interview-hustlers-calendar/internal/apperrors"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/events/dto/responses"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/helpers"
)

func (c *eventsController) ListEvents() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		events, err := c.service.All(ctx)
		if err != nil {
			c.logger.Error(err.Error())
			helpers.Error(rw, req, apperrors.Internal("INTERNAL", err))
			return
		}

		res := responses.ToListEventsResponse(events)

		helpers.Success(rw, req, res, http.StatusOK)
	}
}
