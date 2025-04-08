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