package users

import (
	"net/http"

	"github.com/aevula/interview-hustlers-calendar/internal/logging"
	"github.com/aevula/interview-hustlers-calendar/internal/services/users"
)

type UsersController interface {
	CreateUser(rw http.ResponseWriter, req *http.Request)
}

type usersController struct {
	service users.UsersService
	logger  logging.Logger
}

func NewUsersController(service users.UsersService, logger logging.Logger) UsersController {
	return &usersController{service: service, logger: logger}
}
