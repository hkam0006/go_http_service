package products

import (
	"io"

	repo "github.com/hkam0006/ecom-server/internal/adapters/postgresql/sqlc"
	validate_json "github.com/hkam0006/ecom-server/internal/validator"
	"github.com/jackc/pgx/v5/pgtype"
)

// r.Get("/", h.ListProducts)
// r.Get("/{product_id}", h.GetProductById)
// r.Post("/", h.AddProduct)
//

type Validator interface {
	GetProductById		(id string, format *pgtype.UUID) error
	AddProduct			(body io.ReadCloser, format *repo.CreateProductParams) error
	DeleteProduct		(id string, format *pgtype.UUID) error
}

type validator struct {}

func (v *validator) GetProductById (id string, format *pgtype.UUID) error {
	return format.Scan(id)
}

func (v *validator) AddProduct(body io.ReadCloser, format *repo.CreateProductParams) error {
	return validate_json.NewValidator().Validate(body, format)
}

func (v *validator) DeleteProduct (id string, format *pgtype.UUID) error {
	return format.Scan(id)
}

func NewValidator() Validator {
	return &validator{}
}
