package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/portfolio_be/internal/handler"
	middleware "github.com/vishalpandhare01/portfolio_be/internal/midddleware"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"Message": "server running now !!!",
		})
	})

	app.Post("/adduser", handler.CreateUser)
	app.Get("/allusers", middleware.AuthMiddleware, middleware.AdminRoleMiddleware, handler.GetAllUsers)
	app.Post("/login", handler.LoginUser)

	app.Post("/userprofile", middleware.AuthMiddleware, handler.CreateUserProfile)
	app.Get("/userprofile", middleware.AuthMiddleware, handler.GetUserProfileById)
	app.Delete("/userprofile", middleware.AuthMiddleware, handler.DeleteUserProfileById)
	//

}
