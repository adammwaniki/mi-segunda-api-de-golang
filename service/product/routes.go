package product

import (
	"fmt"
	"net/http"

	"github.com/adammwaniki/mi-segunda-api-de-golang/service/auth"
	"github.com/adammwaniki/mi-segunda-api-de-golang/types"
	"github.com/adammwaniki/mi-segunda-api-de-golang/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct{
	store 		types.ProductStore
	userStore 	types.UserStore
}

func NewHandler(store types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProduct).Methods(http.MethodGet)
	router.HandleFunc("/products/{productID}", h.handleGetProduct).Methods(http.MethodGet)
	router.HandleFunc("/products/create", h.handleCreateProduct).Methods(http.MethodPost) // Create a post method for the products similar to the post users

	// admin routes
	router.HandleFunc("/products", auth.WithJWTAuth(h.handleCreateProduct, h.userStore)).Methods(http.MethodPost)

}

// Handler function to get the slice of products
func (h *Handler) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	// call the GetProducts method
	products, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Write to the user the slice of Products 
	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request){
	// Get JSON payload, it is a product in this case
	var product types.CreateProductPayload
	if err := utils.ParseJSON(r, &product); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate the payload 
	// Similar to user payload validation, this package can validate our structs in ../types/types.go
	if err := utils.Validate.Struct(product); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// Create the product
	err := h.store.CreateProduct(product)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	// Consider implementing a check to see if the product exists
	// If exists, consider prompting the user to increase the quantity of that particular product
	

	utils.WriteJSON(w, http.StatusCreated, product)
}

