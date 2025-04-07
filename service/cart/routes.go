package cart

import (
	"net/http"

	"github.com/adammwaniki/mi-segunda-api-de-golang/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	store 			types.OrderStore
	productStore 	types.ProductStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore) *Handler {
	return &Handler{store: store, productStore: productStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.handleCheckout).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {

}