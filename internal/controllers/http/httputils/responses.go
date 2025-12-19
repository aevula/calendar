package httputils

import (
	"net/http"

	"github.com/aevula/interview-hustlers-calendar/internal/apperrors"

	"github.com/go-chi/render"
)

type EmptyResponse struct{}

type SuccessData struct {
	Data any `json:"data"`
}

type ErrorData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Render[D any, T any](rw http.ResponseWriter, req *http.Request, data D, mapper func(data D) T, status int) {
	render.Status(req, status)
	render.JSON(rw, req, SuccessData{mapper(data)})
}

func Error(rw http.ResponseWriter, req *http.Request, err apperrors.AppError) {
	render.Status(req, err.Status())
	render.JSON(rw, req, ErrorData{Code: err.Code(), Message: err.Error()})
}
