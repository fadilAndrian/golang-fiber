package main

import (
	"learn-fibergo/app/routes"
	"learn-fibergo/database"
)

func main() {
	database.Connect()
	routes.Routes()
}
