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
	app.Delete("/deleteuser/:id", middleware.AuthMiddleware, middleware.AdminRoleMiddleware, handler.DeleteUserByAdmin)
	app.Post("/login", handler.LoginUser)

	app.Post("/userprofile", middleware.AuthMiddleware, handler.CreateUserProfile)
	app.Get("/userprofile/:userName", handler.GetUserProfileByUserName)
	app.Delete("/userprofile", middleware.AuthMiddleware, handler.DeleteUserProfileById)

	app.Post("/addContact/:id", handler.CreateContact)
	app.Get("/getContacts/", middleware.AuthMiddleware, handler.GetUserContacts)
	app.Delete("/deleteContact/:id", middleware.AuthMiddleware, handler.DeleteUserContacts)

}
