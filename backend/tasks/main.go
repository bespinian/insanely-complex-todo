package main

import (
	"log"

	"github.com/bespinian/ict-todo/backend/tasks/internal/handlers"
	"github.com/bespinian/ict-todo/backend/tasks/internal/models"
	"github.com/bespinian/ict-todo/backend/tasks/internal/store"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(requestid.New())
	app.Use(healthcheck.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	api := app.Group("/api")

	store := store.NewMemoryStore()
	store.Add(&models.Task{Name: "Demo Task"})

	taskHandler := handlers.NewTaskHandler(store)

	api.Get("/tasks", taskHandler.List)
	api.Get("/tasks/:id", taskHandler.Get)
	api.Post("/tasks", taskHandler.Create)
	api.Put("/tasks/:id", taskHandler.Update)
	api.Delete("/tasks/:id", taskHandler.Delete)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":3000"))
}
