package route

import (
	"Test/api/controller"
	"Test/database"
	"Test/repository"
	"Test/usecase"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter() *fiber.App {
	//repo
	postRepo := repository.NewPostRepository(database.DB)
	//usecase
	postUseCase := usecase.NewPostUseCase(postRepo)
	//controllers
	postController := controller.NewPostController(postUseCase)

	app := fiber.New()
	v1 := app.Group("api/v1")

	// Serve Swagger UI
	app.Static("/swagger", "./docs")

	// Serve Swagger JSON
	app.Get("/swagger/*", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger.json")
	})

	v1.Post("create", postController.Create)
	v1.Get("all", postController.GetAll)
	v1.Get(":id", postController.GetByID)
	v1.Patch(":id", postController.UpdateByID)
	v1.Delete(":id", postController.DeleteByID)
	return app
}
