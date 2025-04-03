package types

// This will be an interface so that we can easily test it
type UserStore interface {
	GetUserByEmail(email string) (*User, error) // e.g. any instance of the store in the store.go package will be a valid variable for this interface
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

// This will be explained in depth in the routes_test.go package
// We could test the routes associated with the Store using a mock interface of the UserStore
/*
type mockUserStore struct {}

func GetUserByEmail(email string) (*User, error) {
	return nil, nil
}
*/

type User struct {
	ID			int			`json:"id"`
	FirstName	string		`json:"firstName"`
	LastName	string		`json:"lastName"`
	Email		string		`json:"email"`
	Password	string		`json:"-"`
	CreatedAt	string		`json:"createdAt"`
	
}
type RegisterUserPayload struct {
	// Struct definition along with JSON marshalling
	FirstName 	string	`json:"firstName" validate:"required"`
	LastName	string	`json:"lastName" validate:"required"`
	Email		string	`json:"email" validate:"required,email"`
	Password	string	`json:"password" validate:"required,min=3,max=130"`
}

