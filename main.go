package main

import (
	"Go-Projects/api/routes"
	"Go-Projects/command/migration"
)

func main() {
	migration.Execute()
	routes.Routes()
}