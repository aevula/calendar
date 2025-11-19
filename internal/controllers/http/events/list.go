package events

import (
	"net/http"

	"github.com/aevula/interview-hustlers-calendar/internal/apperrors"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/responses"
)

func (c *eventsController) ListEvents() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		events, err := c.service.All(ctx)
		if err != nil {
			c.logger.Error(err.Error())
			responses.Error(rw, req, apperrors.Internal("INTERNAL", err))
			return
		}

		responses.Success(rw, req, events, http.StatusOK)
	}
}
