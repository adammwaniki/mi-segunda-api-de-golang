package main

import (
	"log"
	"os"

	"github.com/adammwaniki/mi-segunda-api-de-golang/config"
	"github.com/adammwaniki/mi-segunda-api-de-golang/db"
	mysqlCfg "github.com/go-sql-driver/mysql" // Rename to prevent naming conflict
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// This main.go is going to help us with testing our database migrations
// It will need to connect to the database
func main() {
	db, err := db.NewMySQLStorage(mysqlCfg.Config{
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

	driver, _ := mysql.WithInstance(db, &mysql.Config{}) // Initialise the database driver
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations", // The source url for where we will store the migrations
		"mysql", // Specify the drivers that we need
		driver, // Pass in the driver itself
	)
	if err != nil {
		log.Fatal(err)
	}

	// We can now call either m.Up or m.Down to either make changes or revert changes respectively
	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}