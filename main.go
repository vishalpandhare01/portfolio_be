package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8000, http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept , Token , Authorization",
	}))
	port := os.Getenv("PORT")
	initializer.ConnectDb()
	migration.SetUpMigration()
	internal.SetUpRoutes(app)
	fmt.Println("🚀 Server running on Port: ", port)
	app.Listen(port)
}
