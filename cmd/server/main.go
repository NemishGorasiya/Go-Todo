// @title Todo API
// @version 1.0
// @description A simple Go Todo REST API
// @host localhost:8080
// @BasePath /
// @schemes http
package main

import (
	"log"
	"net/http"

	"github.com/NemishGorasiya/Go-Todo/internal/db"
	"github.com/NemishGorasiya/Go-Todo/internal/handler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	_ = godotenv.Load()

	db.InitDB()

	r := mux.NewRouter()
	handler.RegisterRoutes(r)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
