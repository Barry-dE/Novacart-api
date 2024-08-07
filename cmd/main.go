package main

import "github.com/Barry-dE/Novacart-api/cmd/api"

func main() {
	
	server := api.NewAPIServer(":5173", nil)
	if err := server.Run(); err != nil{
		panic(err)
	}
	
}