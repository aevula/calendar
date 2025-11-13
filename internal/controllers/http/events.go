package controllers

import (
	"net/http"

	"github.com/aevula/interview-hustlers-calendar/internal/repository"
	"github.com/aevula/interview-hustlers-calendar/internal/services/events"

	"github.com/go-chi/render"
)

func createEvent(app App) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		params := events.CreateEvent{}
		err := render.DecodeJSON(req.Body, &params)
		if err != nil {
			app.Logger().Error(err.Error())
			return
		}

		repo := repository.NewEventRepository(app.Db())

		event, err := events.Create(ctx, repo, params)
		if err != nil {
			app.Logger().Error(err.Error())
			return
		}

		render.JSON(rw, req, event)
	}
}

func updateEvent(app App) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		params := events.UpdateEvent{}
		err := render.DecodeJSON(req.Body, &params)
		if err != nil {
			app.Logger().Error(err.Error())
			return
		}

		repo := repository.NewEventRepository(app.Db())

		event, err := events.Update(ctx, repo, params)
		if err != nil {
			app.Logger().Error(err.Error())
			return
		}

		render.JSON(rw, req, event)
	}
}

func deleteEvent(app App) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		params := events.DeleteEvent{}
		err := render.DecodeJSON(req.Body, &params)
		if err != nil {
			app.Logger().Error(err.Error())
			return
		}

		repo := repository.NewEventRepository(app.Db())

		err = events.Delete(ctx, repo, params)
		if err != nil {
			app.Logger().Error(err.Error())
			return
		}

		render.JSON(rw, req, params)
	}
}

func listEvents(app App) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		params := events.ListEvents{}
		err := render.DecodeJSON(req.Body, &params)
		if err != nil {
			app.Logger().Error(err.Error())
			return
		}

		repo := repository.NewEventRepository(app.Db())

		events, err := events.List(ctx, repo)
		if err != nil {
			app.Logger().Error(err.Error())
			return
		}

		render.JSON(rw, req, events)
	}
}
