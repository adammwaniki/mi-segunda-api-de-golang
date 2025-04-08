package order

import (
	"database/sql"

	"github.com/adammwaniki/mi-segunda-api-de-golang/types"
)

type Store struct {
	db *sql.DB
}

// Repository for the products
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// Insert into orders and get back the id since we need it for the cart
func (s *Store) CreateOrder(order types.Order) (int, error) {
	res, err := s.db.Exec("INSERT INTO orders (userID, total, status, address) VALUES(?, ?, ?, ?)", order.UserID, order.Total, order.Status, order.Address)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId() // needs to target a column with auto_increment e.g. id
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Insert into order_items, we don't need to get back anything
func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items (orderId, productId, quantity, price) VALUES (?, ?, ?, ?)", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	return err
}