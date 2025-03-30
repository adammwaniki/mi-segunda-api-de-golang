package types

type RegisterUserPayload struct {
	// Struct definition along with JSON marshalling
	FirstName 	string	`json:"firstName"`
	LastName	string	`json:"lastName"`
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}