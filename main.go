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
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // Set global body size limit (100 MB)
	})
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in load env: ", err)
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8000, http://localhost:3000 ,http://localhost:3001 ,http://localhost:3002 , https://myportfolioadmin.vercel.app/ , https://nine-portfolio.vercel.app/",
		AllowHeaders: "Origin, Content-Type, Accept , Token , Authorization",
	}))
	port := os.Getenv("PORT")
	if port == "" {
		port = "0.0.0.0:8800"
	}
	initializer.ConnectDb()
	migration.SetUpMigration()
	internal.SetUpRoutes(app)
	fmt.Println("🚀 Server running on Port: ", port)
	app.Listen(port)
}
