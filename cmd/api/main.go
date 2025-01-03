package main

import (
	"log"

	"github.com/Barry-dE/Novacart-api/internal/db"
	"github.com/Barry-dE/Novacart-api/internal/env"
	"github.com/Barry-dE/Novacart-api/internal/store"
	"github.com/lpernett/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Printf("There was an error loading environment variables")
	}

	// Inject configurations
	cfg := Config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:               env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConnections: env.GetInt("DB_MAX_OPEN_CONNECTIONS", 30),
			maxIdleConnections: env.GetInt("DB_MAX_IDLE_CONNECTIONS", 30),
			maxIdleTime:        env.GetString("DB_MAX_IDLE_TIME", "10m"),
		},
	}

	//database store using respository pattern
	db, err := db.NewConnection(cfg.db.addr, cfg.db.maxOpenConnections, cfg.db.maxIdleConnections, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	log.Println("database connection established")

	store := store.NewStorage(db)

	// Initailize application
	app := &application{
		config: cfg,
		store:  store,
	}

	// mount initialized application
	mux := app.mount()

	log.Fatal(app.run(mux))
}
