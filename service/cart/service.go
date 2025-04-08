package cart

import (
	"fmt"

	"github.com/adammwaniki/mi-segunda-api-de-golang/types"
)

// There will be a few business logic decisions handled here

// Function to get a slice of IDs from a slice of cart items
func getCartItemsIDs(items []types.CartItem) ([]int, error) {
	productIds := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 { // doesn't make sense to return zero items
			return nil, fmt.Errorf("invalid quantity for the product %d", item.ProductID)
		}

		productIds[i] = item.ProductID
	}

	return productIds, nil
}

// function to create
// This method will be received by the handler because it needs to access all the repositories
// It will take in the products, the cart items and the userID to make a table association
// It will return the orderId, the total amount for a user to pay and an error
func (h *Handler) CreateOrder(ps []types.Product, item []types.CartItem, userID int) (int, float64, error) {
	// Algo:
	// Check  if all products are actually in stock
	// Calculate the total price
	// Reduce the quantity of products in our database (ACID)
	// Create the order
	// Create the order items
	// Ideally you should wrap this into one SQL transaction. This will be a separate assignment
}