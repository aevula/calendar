package helpers

import (
	"net/http"

	"github.com/aevula/interview-hustlers-calendar/internal/apperrors"

	"github.com/go-chi/render"
)

type SuccessData struct {
	Data any `json:"data"`
}

type ErrorData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Status int
	Body   any
}

func Respond(rw http.ResponseWriter, req *http.Request, res Response) {
	render.Status(req, res.Status)
	render.JSON(rw, req, res.Body)
}

func Success(rw http.ResponseWriter, req *http.Request, data any, status int) {
	render.Status(req, status)
	render.JSON(rw, req, SuccessData{data})
}

func Error(rw http.ResponseWriter, req *http.Request, err apperrors.AppError) {
	render.Status(req, err.Status())
	render.JSON(rw, req, ErrorData{Code: err.Code(), Message: err.Error()})
}
