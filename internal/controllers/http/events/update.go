package events

import (
	"net/http"

	"github.com/aevula/interview-hustlers-calendar/internal/apperrors"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/events/dto/requests"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/responses"

	"github.com/go-chi/render"
)

func (c *eventsController) UpdateEvent() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		params := requests.UpdateEventRequest{}
		err := render.DecodeJSON(req.Body, &params)
		if err != nil {
			responses.Error(rw, req, apperrors.BadRequest("BAD_REQUEST", err))
			return
		}

		cmd := params.ToCommand()

		event, err := c.service.Update(ctx, cmd)
		if err != nil {
			c.logger.Error(err.Error())
			responses.Error(rw, req, apperrors.Internal("INTERNAL", err))
			return
		}

		c.logger.Info("Updated Event", c.logger.Int("ID", int(event.ID)))
		responses.Success(rw, req, event, http.StatusOK)
	}
}
