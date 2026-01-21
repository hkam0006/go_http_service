package orders

import (
	"context"

	repo "github.com/hkam0006/ecom-server/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service interface {
	CreateOrder (context context.Context, user_id pgtype.UUID) (repo.Order, error)
	CreateOrderItems (context context.Context, params repo.CreateOrderItemsParams) ([]repo.OrderItem, error)
	GetProductsByIds (context context.Context, product_ids []pgtype.UUID) ([]repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func (s *svc) CreateOrder (context context.Context, user_id pgtype.UUID) (repo.Order, error) {
	return s.repo.PlaceOrder(context, user_id)
}

func (s *svc) CreateOrderItems (context context.Context, params repo.CreateOrderItemsParams) ([]repo.OrderItem, error) {
	return s.repo.CreateOrderItems(context, params)
}

func (s *svc) GetProductsByIds (context context.Context, product_ids []pgtype.UUID) ([]repo.Product, error) {
	return s.repo.GetProductsByIds(context, product_ids)
}

func NewService(repo repo.Querier) Service {
	return &svc {
		repo: repo,
	}
}
