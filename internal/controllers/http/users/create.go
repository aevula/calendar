package users

import (
	"net/http"

	"github.com/aevula/interview-hustlers-calendar/internal/apperrors"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/responses"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/users/dto/requests"

	"github.com/go-chi/render"
)

func (c *usersController) CreateUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		params := requests.CreateUserRequest{}
		err := render.DecodeJSON(req.Body, &params)
		if err != nil {
			responses.Error(rw, req, apperrors.BadRequest("BAD_REQUEST", err))
			return
		}

		cmd := params.ToCommand()

		user, err := c.service.Create(ctx, cmd)
		if err != nil {
			c.logger.Error(err.Error())
			responses.Error(rw, req, apperrors.Internal("INTERNAL", err))
			return
		}

		c.logger.Info("Created User", c.logger.Int("ID", user.ID))
		responses.Success(rw, req, user, http.StatusCreated)
	}
}
