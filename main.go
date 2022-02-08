package main

import (
	"github.com/pedrosrtavares/go-rest-api/database"
	"github.com/pedrosrtavares/go-rest-api/routes"
)

func main() {
	database.ConnectToDatabase()
	routes.HandleRequest()
}
