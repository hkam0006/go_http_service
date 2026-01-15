package products

import (
	"github.com/go-chi/chi/v5"
)

// Router for /products
func NewRouter() chi.Router {
	r := chi.NewRouter()

	s := NewService()
	h := NewHandler(s)

    r.Get("/", h.ListProducts)

	return r
}
