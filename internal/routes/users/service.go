package users

import (
	"context"

	repo "github.com/hkam0006/ecom-server/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service interface {
	AddUser (ctx context.Context, req repo.AddUserParams) (repo.User, error)
	DeleteUser (ctx context.Context, id pgtype.UUID) (pgtype.UUID, error)
}

type svc struct {
	db repo.Querier
}

func (s *svc) AddUser (ctx context.Context, req repo.AddUserParams) (repo.User, error) {
	return s.db.AddUser(ctx, req)
}

func (s *svc) DeleteUser (ctx context.Context, id pgtype.UUID) (pgtype.UUID, error) {
	return s.db.DeleteUser(ctx, id)
}

func NewService(repo repo.Querier) Service {
	return &svc {
		db: repo,
	}
}
