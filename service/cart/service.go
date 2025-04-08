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
func (h *Handler) createOrder(products []types.Product, cartItems []types.CartItem, userID int) (int, float64, error) {
	// small optimisation prior to the algorithm implementation
	// Since we will require to loop multiple times over the products, we can create a map to speed up lookups
	productsMap := make(map[int]types.Product)
	for _, product := range products {
		productsMap[product.ID] = product
	}

	// Algo:
	// Check  if all products are actually in stock
	// Calculate the total price
	// Reduce the quantity of products in our database // remember to refactor for atomicity(ACID)
	// Create the order
	// Create the order items
	// Ideally you should wrap this into one SQL transaction. This will be a separate assignment

	// stock check
	if err := checkIfCartIsInStock(cartItems, productsMap); err != nil {
		return 0, 0, nil
	}

	// price calculation
	totalPrice := calculateTotalPrice(cartItems, productsMap)

	// reduce quantity of products from the db
	// It would be best to refactor for atomicity e.g. by handling quantities in a separate table the same way orderItems are a separate table
	for _, item := range cartItems {
		product := productsMap[item.ProductID]
		product.Quantity -= item.Quantity

		h.store.UpdateProduct(product)
	}

	// create the order record
	orderID, err := h.orderStore.CreateOrder(types.Order{
		UserID: userID,
		Total: totalPrice,
		Status: "pending", // default
		Address: "some address", // hardcoded for now but it will need a user address table and be able to fetch the default one if a user has multiple addresses e.g. on amazon
	})
	if err != nil {
		return 0, 0, err
	}

	// create order items
	for _, item := range cartItems {
		h.orderStore.CreateOrderItem(types.OrderItem{
			OrderID: orderID,
			ProductID: item.ProductID,
			Quantity: item.Quantity,
			Price: productsMap[item.ProductID].Price,
		})
	}


	return orderID, totalPrice, nil

}

func checkIfCartIsInStock(cartItems []types.CartItem, products map[int]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems{
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d if not available in the store, please refresh your cart", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity requested", product.Name)
		}
	}
	return nil
}

func calculateTotalPrice(cartItems []types.CartItem, products map[int]types.Product) float64 {
	var total float64

	for _, item := range cartItems {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}

	return total
}