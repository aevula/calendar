package users

import (
	"net/http"

	"github.com/aevula/interview-hustlers-calendar/internal/apperrors"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/helpers"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/users/dto/requests"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/users/dto/responses"
)

func (c *usersController) CreateUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		params, err := helpers.ParseBody[requests.CreateUserRequest](req.Body)
		if err != nil {
			helpers.Error(rw, req, apperrors.BadRequest("BAD_REQUEST", err))
			return
		}

		cmd := params.ToCommand()

		user, err := c.service.Create(ctx, cmd)
		if err != nil {
			c.logger.Error(err.Error())
			helpers.Error(rw, req, apperrors.Internal("INTERNAL", err))
			return
		}

		res := responses.ToCreateUserResponse(user)

		c.logger.Info("Created User", c.logger.Int("ID", res.ID))
		helpers.Success(rw, req, res, http.StatusCreated)
	}
}
