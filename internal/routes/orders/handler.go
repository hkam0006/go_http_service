package orders

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	repo "github.com/hkam0006/ecom-server/internal/adapters/postgresql/sqlc"
	json_utils "github.com/hkam0006/ecom-server/internal/json"
	"github.com/jackc/pgx/v5/pgtype"
)

type handler struct {
	service Service
}

// Post JSON: {
// 		user_id,
// 		products: [
// 			{product_id, quantity}
// 		],
// 		discount_code: stirng
// }
func (h *handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var request PlaceOrderRequest

	if err := json_utils.Read(r, &request); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	multiplier := 1.0
	if request.DiscountCode == "TEST_123" {
		multiplier = 0.5
	}
	context := r.Context()

	// Create Order
	order, err := h.service.CreateOrder(context, request.UserID)
	if err != nil {
		log.Println("Error creating order")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	productIDs := make([]pgtype.UUID, 0, len(request.Products))
	qtyByProductID := make(map[pgtype.UUID]int)
	for _, p := range request.Products {
		productIDs = append(productIDs, p.ProductID)
		qtyByProductID[p.ProductID] = int(p.Quantity)
	}

	// Get Products
	products, err := h.service.GetProductsByIds(context, productIDs)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	product_with_price := make([]ProductWithPrice, 0, len(products))
	for _, p := range products {
		item := ProductWithPrice{
			PlaceOrderItemRequest: PlaceOrderItemRequest{
				ProductID: p.ID,
				Quantity:  int32(qtyByProductID[p.ID]),
			},
			PricePerProductInCents: int32(multiplier * float64(p.PriceInCents)),
		}
		product_with_price = append(product_with_price, item)
	}

	productsJSON, err := json.Marshal(product_with_price)
	if err != nil {
		log.Println("Error marshalling products:", err)
		http.Error(w, "Failed to process order items", http.StatusInternalServerError)
		return
	}

	items, err := h.service.CreateOrderItems(context, repo.CreateOrderItemsParams{
		Column1: order.ID,      // order_id
		Column2: productsJSON,  // jsonb array of products {product_id, quantity, price_in_cents}
	})

	if err != nil {
		log.Println("Error creating order items:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := struct {
		Order    repo.Order        `json:"order"`
		Products []repo.OrderItem  `json:"products"`
	}{
		Order:    order,
		Products: items,
	}

	json_utils.Write(w, http.StatusCreated, res)
}

func (h *handler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	order_id := chi.URLParam(r, "order_id")

	var order_uuid pgtype.UUID

	if err := order_uuid.Scan(order_id); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := h.service.GetOrderById(r.Context(), order_uuid)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json_utils.Write(w, http.StatusOK, order)
}


func NewHandler(s Service) *handler{
	return &handler{
		service: s,
	}
}
