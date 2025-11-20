package users

import (
	"net/http"

	"github.com/aevula/interview-hustlers-calendar/internal/apperrors"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/httputils"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/users/dto/requests"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/users/dto/responses"
)

func (c *usersController) CreateUser(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	params, err := httputils.ParseBody(req.Body, requests.CreateUserRequest{})
	if err != nil {
		httputils.Error(rw, req, apperrors.BadRequest("BAD_REQUEST", err))
		return
	}

	cmd := params.ToCommand()

	user, err := c.service.Create(ctx, cmd)
	if err != nil {
		c.logger.Error(err.Error())
		httputils.Error(rw, req, apperrors.Internal("INTERNAL", err))
		return
	}

	c.logger.Info("Created User", c.logger.Int("ID", int(user.ID)))
	httputils.Render(rw, req, user, responses.ToCreateUserResponse, http.StatusCreated)
}
