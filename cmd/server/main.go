// @title Todo API
// @version 1.0
// @description A simple Go Todo REST API
// @host localhost:8080
// @BasePath /
// @schemes http
package main

import (
	_ "github.com/NemishGorasiya/Go-Todo/docs"

	"log"
	"net/http"

	"github.com/NemishGorasiya/Go-Todo/internal/db"
	"github.com/NemishGorasiya/Go-Todo/internal/handler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	db.InitDB()

	r := mux.NewRouter()
	handler.RegisterRoutes(r)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
