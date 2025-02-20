package main

import (
	"golang-mvc/app"
	"net"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

// @title TodoAPI
// @version 1.0
// @description This is a todo API.
// @host localhost:1006
// @BasePath /api/v1
// @SecurityDefinitions.Apikey JWT
// @In header
// @Name Authorization

func main() {
	err := godotenv.Load()
	
	if err != nil {
		panic(err)
	}

	route := app.Init()

	port := os.Getenv("PORT")

	ln, err := net.Listen("tcp", "localhost:" + port)

	if err != nil {
		panic(err)
	}

	_ = http.Serve(ln, route)
}