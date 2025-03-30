package user

import (
	"net/http"

	"github.com/adammwaniki/mi-segunda-api-de-golang/types"
	"github.com/adammwaniki/mi-segunda-api-de-golang/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store *types.UserStore
}

func NewHandler() *Handler{
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request){

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request){
	// Get JSON payload,
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// Check if user exists

	// If user does not exist, create a new user
}