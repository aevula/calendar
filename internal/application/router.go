package application

import "github.com/go-chi/chi/v5"

func newRouter() *chi.Mux {
	return chi.NewRouter()
}
