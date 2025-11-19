package events

import (
	"net/http"

	"github.com/aevula/interview-hustlers-calendar/internal/apperrors"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/events/dto/requests"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/events/dto/responses"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/helpers"
)

func (c *eventsController) CreateEvent() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		params, err := helpers.ParseBody[requests.CreateEventRequest](req.Body)
		if err != nil {
			helpers.Error(rw, req, apperrors.BadRequest("BAD_REQUEST", err))
			return
		}

		cmd, err := params.ToCommand()
		if err != nil {
			helpers.Error(rw, req, apperrors.BadRequest("BAD_REQUEST", err))
			return
		}

		event, err := c.service.Create(ctx, cmd)
		if err != nil {
			c.logger.Error(err.Error())
			helpers.Error(rw, req, apperrors.Internal("INTERNAL", err))
			return
		}

		res := responses.ToCreateEventResponse(event)

		c.logger.Info("Created Event", c.logger.Int("ID", int(event.ID)))
		helpers.Success(rw, req, res, http.StatusCreated)
	}
}
