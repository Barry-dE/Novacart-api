package main

import (
	"github.com/Barry-dE/REST-API-ECS/cmd/api"
)

func main() {
	server := api.NewAPIServer(":5173", nil)
	if err := server.Run(); err != nil{
		panic(err)
	}
	
}