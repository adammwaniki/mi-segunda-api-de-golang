package user

import (
	"fmt"
	"net/http"

	"github.com/adammwaniki/mi-segunda-api-de-golang/service/auth"
	"github.com/adammwaniki/mi-segunda-api-de-golang/types"
	"github.com/adammwaniki/mi-segunda-api-de-golang/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler{
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request){
	// Get JSON payload, just like the RegisterUserPayload
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate the payload, just like handleRegister user route
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return 
	}

	// Check if the hashed password in the db and the inputted password match
	if !auth.ComparePasswords(u.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	// Passing in an empty token for now prior to implementation of JWT
	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": ""})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request){
	// Get JSON payload,
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate the payload
	// Option 1: write a function that checks the payload.email, payload.username etc. but can become verbose if there are many endpoints
	// Option 2: use a validator package e.g., github.com/go-playground/validator/v10
	// This package can validate our structs in ../types/types.go
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// Check if user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// If user does not exist, create a new user
	err = h.store.CreateUser(types.User{
		FirstName: 	payload.FirstName,
		LastName:  	payload.LastName,
		Email: 		payload.Email,
		Password: 	hashedPassword, // Updated to the hashed password
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}