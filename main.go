package main

import (
	"authJWT/config"
	"authJWT/routes"
)

func main() {
	// connect to database
	config.SetupDatabase()

	// initialize the router
	r := routes.SetupRouter()

	// start the server
	r.Run(":8080")
}
