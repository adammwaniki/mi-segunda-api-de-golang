package main

import (
	"database/sql"
	"log"

	"github.com/adammwaniki/mi-segunda-api-de-golang/cmd/api"
	"github.com/adammwaniki/mi-segunda-api-de-golang/config"
	"github.com/adammwaniki/mi-segunda-api-de-golang/db"
	"github.com/go-sql-driver/mysql"
)

// This will be the entrypoint for our API

func main(){
	// Creating a new database connection
	// This part is usually in your environment variables and not hardcoded
	db, err := db.NewMySQLStorage(mysql.Config{
		User: 					config.Envs.DBUser,
		Passwd: 				config.Envs.DBPassword, // Can be anything
		Addr: 					config.Envs.DBAddress, // Is a combination of the host and the port
		DBName: 				config.Envs.DBName, // Can be anything related to your project
		Net: 					"tcp", // tcp is the default but it can be tcp6 or even unix
		AllowNativePasswords: 	true, // Allow the native password authentication method
		ParseTime: 				true, // Parses time values to time.Time
	})
	if err != nil {
		log.Fatal(err)
	}

	// Call the initStorage function and remember to pass in the db
	// NB: Remember to ceate the database first e.g. in your sql editor
	initStorage(db)

	// Creating a new API/server instance
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// Initialise the database connection
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected")
}