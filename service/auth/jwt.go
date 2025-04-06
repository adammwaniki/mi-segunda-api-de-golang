package auth

import (
	"strconv"
	"time"

	"github.com/adammwaniki/mi-segunda-api-de-golang/config"
	"github.com/golang-jwt/jwt/v5"
)


func CreateJWT(secret []byte, userID int) (string, error) {
	// Creating a variable to hold the expiration time of the token
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds) // JWTExpirationInSeconds gotten from the env

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": strconv.Itoa(userID),
		"expeiredAt": time.Now().Add(expiration).Unix(), // This will be the expiration of the token e.g. the expiration of the server in this case
	})

	tokenString, err := token.SignedString(secret) // make a signed token
	if err != nil {
		return "", err
	}
	return tokenString, nil
}