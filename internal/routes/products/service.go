package products

import (
	"context"

	repo "github.com/hkam0006/ecom-server/internal/adapters/postgresql/sqlc"
)

// Interacting with database
type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	GetProduct(ctx context.Context, id int64) (repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) GetProduct(ctx context.Context, id int64) (repo.Product, error) {
	return s.repo.FindProductsByID(ctx, id)
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}
