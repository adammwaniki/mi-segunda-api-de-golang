package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/adammwaniki/mi-segunda-api-de-golang/config"
	"github.com/adammwaniki/mi-segunda-api-de-golang/types"
	"github.com/adammwaniki/mi-segunda-api-de-golang/utils"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserKey contextKey = "userID"


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

// Authentication wrapper function
func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Start by checking if user has a token in the headers or cookies or wherever
		// Validate the JWT
		// If valid, fetch userID from the DB (since we get userID from the token)
		// set context "userID"

		// get the token from the  user request
		tokenString := utils.GetTokenFromRequest(r)

		// validate the token
		token, err := validateJWT(tokenString)
		if err != nil {
			log.Printf("failed to validate token: %v", err)
			permissionDenied(w)
			return
		}
		
		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}

		// fetch userID from the token
		// first we obtain the id from the token
		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)

		userID, _ := strconv.Atoi(str)

		u, err := store.GetUserByID(userID)
		// if user does't exist
		if err != nil {
			log.Printf("failedd to get  user by id: %v", err)
			permissionDenied(w)
			return
		}

		// finally add user to the context
		ctx := r.Context() // get the context from the request
		ctx = context.WithValue(ctx, UserKey, u.ID) // initialise the context // the compiler may complain about inputting "userID" as a string directly so we shall make a type for it
		r = r.WithContext(ctx) // modify the context of the request

		handlerFunc(w, r)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // here we cast the Method
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}

// This will be used a lot so we can just make it a function
func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

// Helper func to get the userID from the context
func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int) //casting to integer because we know it is an integer to prevent compiler complaints
	if !ok {
		return -1
	}

	return userID
}



/*

	*/