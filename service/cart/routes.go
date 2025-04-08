package cart

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
	store      types.ProductStore
	orderStore types.OrderStore
	userStore  types.UserStore
}

func NewHandler(
	store types.ProductStore,
	orderStore types.OrderStore,
	userStore types.UserStore,
) *Handler {
	return &Handler{
		store:      store,
		orderStore: orderStore,
		userStore:  userStore,
	}
}

// We need to add authentication
// We can use a higher order function aka a decorated pattern
// The idea is to wrap our func below in an authentication handler
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", auth.WithJWTAuth(h.handleCheckout, h.userStore)).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	// We expect to receive some sort of cart object from the frontend
	// e.g. an array of items and their quantity
	// Start by parsing this json from the frontend
	userID := auth.GetUserIDFromContext(r.Context()) // This will come from the JWT once we implement authentication
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

	// Get products
	// We need these products to be handled by id as a slice of ids. This can be handled in the service file 
	productIDs, err := getCartItemsIDs(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	ps, err := h.store.GetProductsByIDs(productIDs)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	orderID, totalPrice, err := h.createOrder(ps, cart.Items, userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"total_price": totalPrice,
		"order_id": orderID,
	})

}