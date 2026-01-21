package orders

import (
	"github.com/jackc/pgx/v5/pgtype"
)


type PlaceOrderItemRequest struct {
	ProductID 	pgtype.UUID 	`json:"product_id"`
	Quantity  	int32    			`json:"quantity"`
}

type ProductWithPrice struct {
	PlaceOrderItemRequest
	PricePerProductInCents int32   `json:"price_per_product_in_cents"`
}

type PlaceOrderRequest struct {
	UserID       pgtype.UUID           		`json:"user_id"`
	Products     []PlaceOrderItemRequest 	`json:"products"`
	DiscountCode string           			`json:"discount_code"`
}
