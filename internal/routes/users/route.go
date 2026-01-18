package users

import (
	"github.com/go-chi/chi/v5"
	repo "github.com/hkam0006/ecom-server/internal/adapters/postgresql/sqlc"
)

func NewRouter(db repo.DBTX) chi.Router {
	r := chi.NewRouter()

	v := NewValidator()
	s := NewService(repo.New(db))
	h := NewHandler(s, v)

	r.Post("/", h.AddUser)
	r.Delete("/{user_id}", h.DeleteUser)

	return r
}
