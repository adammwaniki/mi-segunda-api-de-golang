package product

import (
	"net/http"

	"github.com/adammwaniki/mi-segunda-api-de-golang/types"
	"github.com/adammwaniki/mi-segunda-api-de-golang/utils"
	"github.com/gorilla/mux"
)

type Handler struct{
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodGet)
}

// Handler function to get the slice of products
func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	// call the method
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Write to the user the slice of Products 
	utils.WriteJSON(w, http.StatusOK, ps)
}