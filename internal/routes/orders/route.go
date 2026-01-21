package orders

import (
	"github.com/go-chi/chi/v5"
	repo "github.com/hkam0006/ecom-server/internal/adapters/postgresql/sqlc"
)

func NewRouter(db repo.DBTX) chi.Router {
	r := chi.NewRouter()

	s := NewService(repo.New(db))
	h := NewHandler(s)

	r.Post("/", h.CreateOrder)

	return r
}
