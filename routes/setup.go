package routes

import "github.com/gofiber/fiber/v2"

func SetupUserRoutes(app *fiber.App) {
	usersRoute := app.Group("/api/users")

	usersRoute.Get("/", GetUsers)
	usersRoute.Post("/", CreateUser)
	usersRoute.Get("/:id", GetUserByID)
	usersRoute.Get("/email/:email", GetUserByEmail)
	usersRoute.Delete("/:id", DeleteUser)
	usersRoute.Put("/:id", UpdateUser)
}

func SetupPermissionRoutes(app *fiber.App) {
	permRoutes := app.Group("/api/permissions")

	permRoutes.Get("/", GetPermissions)
	permRoutes.Post("/", CreatePermission)
	permRoutes.Get("/:id", GetPermission)
	permRoutes.Delete("/:id", DeletePermission)
	permRoutes.Put("/:id", UpdatePermission)
}

func SetupRoutes(app *fiber.App) {
	SetupUserRoutes(app)
	SetupPermissionRoutes(app)
}
