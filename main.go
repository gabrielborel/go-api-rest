package main

import (
	"api-rest/database"
	"api-rest/routes"
	"fmt"
)

func main() {
	database.ConnectWithDatabase()

	fmt.Println("Iniciando servidor")
	routes.HandleRequest()
}
