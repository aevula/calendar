package helpers

import (
	"io"

	"github.com/go-chi/render"
)

func ParseBody[T any](body io.ReadCloser) (T, error) {
	var params T
	err := render.DecodeJSON(body, &params)
	return params, err
}
