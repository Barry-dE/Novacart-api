package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Barry-dE/Novacart-api/cmd/api"
	"github.com/Barry-dE/Novacart-api/config"
	"github.com/Barry-dE/Novacart-api/database"
	"github.com/go-sql-driver/mysql"
)

func main() {
	//create new database instance
	db, err := database.CreateNewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAdrress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: false,
		ParseTime:            true,
	})
	fmt.Println()
	if err != nil {
		panic(err)
	}

	initDatabase(db)

	//Create API server
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		panic(err)
	}
}

// connect to the database
func initDatabase(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	log.Println("Connection to database was successful")
}
