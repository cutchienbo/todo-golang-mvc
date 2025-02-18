package main

import (
	"golang-mvc/app"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
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

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Cho phép tất cả các origin
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}), // Các phương thức HTTP được phép
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // Các header cho phép
	)(route)

	http.Handle("/", corsHandler)

	port := os.Getenv("PORT")

	ln, err := net.Listen("tcp", "localhost:" + port)

	if err != nil {
		panic(err)
	}

	_ = http.Serve(ln, route)
}