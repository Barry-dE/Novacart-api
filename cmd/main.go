package main

import (
	"database/sql"
	"log"

	"github.com/Barry-dE/Novacart-api/cmd/api"
	"github.com/Barry-dE/Novacart-api/config"
	"github.com/Barry-dE/Novacart-api/database"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err :=  database.CreateNewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAdrress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	}); 

	if err !=nil {
		panic(err)
	}

	server := api.NewAPIServer(":5173", db )
	if err := server.Run(); err != nil{
		panic(err)
	}

	initDatabase(db)
	
}

func initDatabase(db *sql.DB){
if err := db.Ping(); err != nil{
	panic(err)
}

log.Println("Connection to database was successful")

}