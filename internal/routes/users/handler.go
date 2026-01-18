package users

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	repo "github.com/hkam0006/ecom-server/internal/adapters/postgresql/sqlc"
	"github.com/hkam0006/ecom-server/internal/json"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
)

type handler struct {
	service Service
	validator Validator
}

func NewHandler (s Service, v Validator) *handler {
	return &handler {
		service: s,
		validator: v,
	}
}

func (h *handler) AddUser (w http.ResponseWriter, r *http.Request) {
	var req repo.AddUserParams

	if err := h.validator.AddUser(r.Body, &req); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	req.Password = string(hashedPassword)

	user, err := h.service.AddUser(r.Context(), req)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, user)
}

func (h *handler) DeleteUser (w http.ResponseWriter, r *http.Request) {
	user_id := chi.URLParam(r, "user_id")

	var uuid pgtype.UUID

	if err := h.validator.DeleteUser(user_id, &uuid); err != nil {
		log.Println("Invalid UUID")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.DeleteUser(r.Context(), uuid)

	if err != nil {
		log.Println("Error deleting user")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, id)
}
