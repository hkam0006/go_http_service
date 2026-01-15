package products

import (
	"log"
	"net/http"

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
