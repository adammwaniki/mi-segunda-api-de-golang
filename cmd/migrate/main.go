package main

import (
	"log"

	"github.com/adammwaniki/mi-segunda-api-de-golang/cmd/api"
	"github.com/adammwaniki/mi-segunda-api-de-golang/db"
	"github.com/go-sql-driver/mysql"
)

// This will be the entrypoint for our API

func main(){
	// Creating a new database connection
	// This part is usually in your environment variables and not hardcoded
	db, err := db.NewMySQLStorage(mysql.Config{
		User: "root",
		Passwd: "asd", // Can be anything
		Addr: "127.0.1:3306", // Is a combination of the host and the port
		DBName: "ecommerce", // Can be anything related to your project
		Net: "tcp", // tcp is the default but it can be tcp6 or even unix
		AllowNativePasswords: true, // Allow the native password authentication method
		ParseTime: true, // Parses time values to time.Time
	})

	// Creating a new API/server instance
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}