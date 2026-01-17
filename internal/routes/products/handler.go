package products

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hkam0006/ecom-server/internal/json"
)

// Handle business logic here.
type handler struct {
	service Service // dependency
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// 1. Call Service -> list products
	products, err := h.service.ListProducts(r.Context())

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. Return JSON in a HTTP response
	json.Write(w, http.StatusOK, products)
}

func (h *handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	product_id_str := chi.URLParam(r, "product_id")

	id, conv_err := strconv.ParseInt(product_id_str, 10, 64)

	if conv_err != nil {
		log.Println("Converting error")
		http.Error(w, conv_err.Error(), http.StatusInternalServerError)
		return
	}

	product, err := h.service.GetProduct(ctx, id)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusAccepted, product)

}
