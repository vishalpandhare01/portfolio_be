package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/vishalpandhare01/portfolio_be/initializer"
	"github.com/vishalpandhare01/portfolio_be/internal"
	"github.com/vishalpandhare01/portfolio_be/migration"
)

func main() {
	app := fiber.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in load env: ", err)
	}
	port := os.Getenv("PORT")
	initializer.ConnectDb()
	migration.SetUpMigration()
	internal.SetUpRoutes(app)
	fmt.Println("🚀 Server running on Port: ", port)
	app.Listen("0.0.0.0:" + port)
}
