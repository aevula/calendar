package httputils

import (
	"io"

	"github.com/go-chi/render"
)

func ParseBody[T any](body io.ReadCloser, params T) (T, error) {
	err := render.DecodeJSON(body, &params)
	return params, err
}
