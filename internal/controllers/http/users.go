package controllers

import (
	"net/http"

	"github.com/aevula/interview-hustlers-calendar/internal/repository"
	"github.com/aevula/interview-hustlers-calendar/internal/services/users"

	"github.com/go-chi/render"
)

func createUser(app App) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		params := users.CreateUser{}
		err := render.DecodeJSON(req.Body, &params)
		if err != nil {
			app.Logger().Error(err.Error())
			return
		}

		repo := repository.NewUserRepository(app.Db())

		user, err := users.Create(ctx, repo, params)
		if err != nil {
			app.Logger().Error(err.Error())
			return
		}

		render.JSON(rw, req, user)
	}
}
