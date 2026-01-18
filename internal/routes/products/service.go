package products

import (
	"context"

	repo "github.com/hkam0006/ecom-server/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

// Interacting with database
type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	GetProduct(ctx context.Context, id pgtype.UUID) (repo.Product, error)
	CreateProduct(ctx context.Context, params repo.CreateProductParams) (repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) GetProduct(ctx context.Context, id pgtype.UUID) (repo.Product, error) {
	return s.repo.FindProductsByID(ctx, id)
}

func (s *svc) CreateProduct(ctx context.Context, params repo.CreateProductParams) (repo.Product, error) {
	return s.repo.CreateProduct(ctx, params)
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}
