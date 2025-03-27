package config

import (
	"fmt"
	"os"
)

type Config struct {
	PublicHost 	string
	Port 		string
	DBUser		string
	DBPassword	string
	DBAddress	string
	DBName		string
}

// Declaring a global variable that allows us to skip initialising the initConfig function every time
var Envs = initConfig()


// This function will return the configuration object
func initConfig() Config {
	// return the struct
	// We can get the details using a getter function e.g., getEnv
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port: 		getEnv("PORT", "8080"),
		DBUser: 	getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "mypassword"),
		DBAddress: 	fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")), // Dynamically added
		DBName: getEnv("DB_NAME", "ecommerce"),
		
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

