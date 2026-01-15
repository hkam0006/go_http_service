package products

import (
	"github.com/go-chi/chi/v5"
	repo "github.com/hkam0006/ecom-server/internal/adapters/postgresql/sqlc"
)

// Router for /products
func NewRouter(db repo.DBTX) chi.Router {
	r := chi.NewRouter()

	s := NewService(repo.New(db))
	h := NewHandler(s)

    r.Get("/", h.ListProducts)

	return r
}
