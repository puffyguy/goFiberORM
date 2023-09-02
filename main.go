package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
	"github.com/puffyguy/goFiberORM/controllers"
	"github.com/puffyguy/goFiberORM/database"
	"github.com/puffyguy/goFiberORM/models"
)

func Init() {
	models.DB = database.GetMySQLDB()
	// models.DB = database.GetMongoDB()
	fmt.Println("Successfully connected to DB...")
}

func main() {
	Init()
	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")
	setupRoutes(v1)
	log.Fatal(app.Listen("localhost:8081"))
}

/*
# setupRoutes - Initializing endpoints
- every endpoint have to be suffixed with "/api/v1/"
*/
func setupRoutes(v1 fiber.Router) {
	v1.Get("/", controllers.Welcome)
	v1.Get("/books", controllers.GetBooks)
	v1.Get("/books/:isbn", controllers.GetBook)
	v1.Post("/books", controllers.NewBook)
	v1.Put("/books/:isbn", controllers.UpdateBook)
	v1.Delete("/books/:isbn", controllers.DeleteBook)
}
