package main

import (
	"catalog/api/controllers"
)

func main() {
	server := controllers.Server{}
	server.Init()
	defer server.CloseDB()

	server.Serve(":8080")
}
