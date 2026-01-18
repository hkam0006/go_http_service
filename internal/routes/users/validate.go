package users

import (
	"io"

	repo "github.com/hkam0006/ecom-server/internal/adapters/postgresql/sqlc"
	validate_json "github.com/hkam0006/ecom-server/internal/validator"
	"github.com/jackc/pgx/v5/pgtype"
)

type Validator interface {
	AddUser 		(body io.ReadCloser, format *repo.AddUserParams) 	error
	DeleteUser 		(id string, format *pgtype.UUID) 					error
}

type validator struct {}

func (v *validator) AddUser (body io.ReadCloser, format *repo.AddUserParams) error {
	if err := validate_json.NewValidator().Validate(body, format); err != nil {
		return err
	}

	return nil
}

func (v *validator) DeleteUser (id string, format *pgtype.UUID) error {
	return format.Scan(id)
}

func NewValidator() Validator {
	return &validator{}
}
