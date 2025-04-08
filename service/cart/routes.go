package cart

import (
	"fmt"
	"net/http"

	"github.com/adammwaniki/mi-segunda-api-de-golang/types"
	"github.com/adammwaniki/mi-segunda-api-de-golang/utils"
	"github.com/go-playground/validator/v10"
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
	// We expect to receive some sort of cart object from the frontend
	// e.g. an array of items and their quantity
	// Start by parsing this json from the frontend
	var cart types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validating the cart just like the other endpoints now that we have parsed it
	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// We now need to loop over the cart items and perform validation on the items themselves
	// e.g. they may be out of stock or don't exist etc.
	// We can create a service file in the cart directory to help handle business logic

	// get products
	ps, err := h.productStore.GetProducts(productIDs)

}