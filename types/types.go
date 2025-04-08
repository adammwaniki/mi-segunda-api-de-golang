package types

import "time"

type User struct {
	ID			int			`json:"id"`
	FirstName	string		`json:"firstName"`
	LastName	string		`json:"lastName"`
	Email		string		`json:"email"`
	Password	string		`json:"-"`
	CreatedAt	string		`json:"createdAt"`
	
}

type Order struct {
	ID			int			`json:"id"`
	UserID		int			`json:"userID"`
	Total		float64		`json:"total"`
	Status 		string		`json:"status"`
	Address		string		`json:"address"`
	CreatedAt	time.Time	`json:"createdAt"`
}

type OrderItem struct {
	ID			int			`json:"id"`
	OrderID		int			`json:"orderID"`
	ProductID	int			`json:"productID"`
	Quantity 	int			`json:"quatity"`
	Price		float64		`json:"price"`
	CreatedAt	time.Time	`json:"createdAt"`
}

type Product struct {
	ID			int			`json:"id"` // Look into handling product IDs and user IDs using UUID or other means to prevent collisions
	Name 		string		`json:"name"` // Consider adding another field for product identifiers (SKU --internal use and/or UPC (Universal Product Code) --external use) to help handle getProductBySKU/getProductByUPC
	Description	string		`json:"description"`
	Image		string		`json:"image"`
	Price		float64		`json:"price"`
	Quantity	int			`json:"quantity"` // This is not the best way to handle the quantity since it is not atomic (ACID) hence with multiple concurrent requests the reported quantity may be a false value
	CreatedAt	time.Time	`json:"createdAt"`
}

type CartItem struct {
	ProductID 	int `json:"productID"`
	Quantity	int `json:"quantity"`
}

// This will be an interface so that we can easily test it
type UserStore interface {
	GetUserByEmail(email string) (*User, error) // e.g. any instance of the store in the store.go package will be a valid variable for this interface
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

// Interface for the repository of orders in cart
// The cart just receives Orders hence we don't need a CartStore
type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
}

// interface for the repository of products just like the userstore and orderstore above
type ProductStore interface {
	GetProducts() ([]*Product, error) // function that returns a slice of all the Products
	GetProductByID(id int) (*Product, error)
	GetProductsByIDs(ids []int) ([]Product, error) // method to help get products that are in the cart. We pass in a slice of ids and get back a slice of the products
	GetProductByName(name string) (*Product, error)
	CreateProduct(CreateProductPayload) error
	UpdateProduct(Product) error
}

type CreateProductPayload struct{
	Name		string	`json:"name" validate:"required"`
	Description	string	`json:"description" validate:"required"`
	Image		string	`json:"image" validate:"required"`
	Price		float64	`json:"price" validate:"required"` 
	Quantity	int		`json:"quantity" validate:"required"` // Remember to find a better way to handle quantity
}

type RegisterUserPayload struct {
	// Struct definition along with JSON marshalling
	FirstName 	string	`json:"firstName" validate:"required"`
	LastName	string	`json:"lastName" validate:"required"`
	Email		string	`json:"email" validate:"required,email"`
	Password	string	`json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	// Struct definition along with JSON marshalling
	Email		string	`json:"email" validate:"required,email"`
	Password	string	`json:"password" validate:"required"`
}

type CartCheckoutPayload struct {
	Items	[]CartItem	`json:"items" validate:"required"`
}