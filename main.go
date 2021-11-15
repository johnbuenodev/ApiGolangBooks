package main

import (
	"apiGolang/database"
	"apiGolang/server"
)

func main() {

	//Antes de rodar o servidor rodar o banco de dados/migrations
	database.StartDB()

	server := server.NewServer()

	server.Run()

}

//docker-compose up --build

//go run main.go

//go run database/seeder/main.go
