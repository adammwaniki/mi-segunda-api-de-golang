package product

import (
	"fmt"
	"net/http"

	"github.com/adammwaniki/mi-segunda-api-de-golang/types"
	"github.com/adammwaniki/mi-segunda-api-de-golang/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct{
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProduct).Methods(http.MethodGet)
	router.HandleFunc("/products/create", h.handleCreateProduct).Methods(http.MethodPost) // Create a post method for the products similar to the post users

}

// Handler function to get the slice of products
func (h *Handler) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	// call the GetProducts method
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Write to the user the slice of Products 
	utils.WriteJSON(w, http.StatusOK, ps)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request){
	// Get JSON payload,
	var payload types.RegisterProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate the payload 
	// Similar to user payload validation, this package can validate our structs in ../types/types.go
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// Check if product exists
	_, err := h.store.GetProductByName(payload.Name)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("product with name %s already exists", payload.Name))
		return
	}

	// If product does not exist, create a new product
	err = h.store.CreateProduct(types.Product{
		Name: 			payload.Name,
		Description:  	payload.Description,
		Image: 			payload.Image,
		Price: 			payload.Price, 
		Quantity: 		payload.Quantity,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

