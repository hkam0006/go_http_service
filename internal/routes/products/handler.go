package products

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	repo "github.com/hkam0006/ecom-server/internal/adapters/postgresql/sqlc"
	response "github.com/hkam0006/ecom-server/internal/json"
	"github.com/jackc/pgx/v5/pgtype"
)

// Handle business logic here.
type handler struct {
	service Service // dependency
	validator Validator
}

func NewHandler(service Service, validator Validator) *handler {
	return &handler{
		service: service,
		validator: validator,
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
	response.Write(w, http.StatusOK, products)
}

func (h *handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "product_id")

	var pgUUID pgtype.UUID

	if err := pgUUID.Scan(uuid); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := h.service.GetProduct(r.Context(), pgUUID)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Write(w, http.StatusOK, product)
}

func (h *handler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var req repo.CreateProductParams

	if err := h.validator.Validate(r.Body, &req); err != nil {
		log.Println(("Invalid Request Body"))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := h.service.CreateProduct(r.Context(), req)

	if err != nil {
		log.Println("Error ")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Write(w, http.StatusCreated, product)
}
