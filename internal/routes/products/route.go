package products

import (
	"github.com/go-chi/chi/v5"
	repo "github.com/hkam0006/ecom-server/internal/adapters/postgresql/sqlc"
)

// Router for /products
func NewRouter(db repo.DBTX) chi.Router {
	r := chi.NewRouter()

	v := NewValidator()
	s := NewService(repo.New(db))
	h := NewHandler(s, v)

    r.Get("/", h.ListProducts)
    r.Get("/{product_id}", h.GetProductById)
    r.Post("/", h.AddProduct)

	return r
}
