package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

type Config struct {
	PublicHost 				string
	Port 					string
	DBUser					string
	DBPassword				string
	DBAddress				string
	DBName					string
	JWTExpirationInSeconds 	int64
	JWTSecret				string
}

// Declaring a global variable that allows us to skip initialising the initConfig function every time
var Envs = initConfig()


// This function will return the configuration object
func initConfig() Config {
	// Allowing our program to call environment variables from a .env file
	godotenv.Load()
	
	// return the struct
	// We can get the details using a getter function e.g., getEnv
	return Config{
		PublicHost: 			getEnv("PUBLIC_HOST", "http://localhost"),
		Port: 					getEnv("PORT", "8080"),
		DBUser: 				getEnv("DB_USER", "root"),
		DBPassword: 			getEnv("DB_PASSWORD", "root"), // The password that I set to access mysql on localhost
		DBAddress: 				fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")), // Dynamically added
		DBName: 				getEnv("DB_NAME", "ecommerce"),
		JWTSecret: 				getEnv("JWT_SECRET", "not-secret-secret-anymore?"), // this secret will be passed in during login
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7), // This token will be valid for a week by default
	}
}

// takes in a key and a fallback value in case the key does not exist
// returns a value
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

// Helper function to help get environment variables that are integers since the one above handles only strings
func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64) // helps parse the integer
		if err != nil {
			return fallback
		}

		return i
	}
	return fallback
}
