package main

import (
	"log"

	"github.com/Barry-dE/Novacart-api/internal/env"
)

func main() {
	cfg := Config{
		addr: env.GetString("ADDR",":8080")
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
