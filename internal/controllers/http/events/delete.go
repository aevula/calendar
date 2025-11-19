package events

import (
	"net/http"

	"github.com/aevula/interview-hustlers-calendar/internal/apperrors"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/events/dto/requests"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/events/dto/responses"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/helpers"
)

func (c *eventsController) DeleteEvent() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		params, err := helpers.ParseBody[requests.DeleteEventRequest](req.Body)
		if err != nil {
			helpers.Error(rw, req, apperrors.BadRequest("BAD_REQUEST", err))
			return
		}

		cmd := params.ToCommand()

		err = c.service.Delete(ctx, cmd)
		if err != nil {
			c.logger.Error(err.Error())
			helpers.Error(rw, req, apperrors.Internal("INTERNAL", err))
			return
		}

		res := responses.DeleteEventResponse{}

		c.logger.Info("Deleted Event", c.logger.Int("ID", cmd.ID))
		helpers.Success(rw, req, res, http.StatusNoContent)
	}
}
