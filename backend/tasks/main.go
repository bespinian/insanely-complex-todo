package main

import (
	"log"
	"os"

	"github.com/bespinian/ict-todo/backend/tasks/internal"
	"github.com/bespinian/ict-todo/backend/tasks/internal/events"
	"github.com/bespinian/ict-todo/backend/tasks/internal/handlers"
	"github.com/bespinian/ict-todo/backend/tasks/internal/store"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New(fiber.Config{ErrorHandler: handlers.ErrorHandler})
	app.Use(logger.New())
	app.Use(requestid.New())
	app.Use(healthcheck.New())

	taskStore := configureTaskStore()
	configureEventQueues()
	configureAPIHandler(app, taskStore)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":3000"))
}

func configureTaskStore() internal.TaskStore {
	if mongoUri := os.Getenv("MONGODB_URI"); mongoUri != "" {
		log.Print("MONGODB_URI is set. Using MongoDB as storage.")
		db, err := store.NewMongoDatabase(mongoUri, "ict")
		if err != nil {
			log.Printf("Could not connect to MongoDB: %s. Falling back to in-memory store.", err)
			return store.NewMemoryStore()
		} else {
			return store.NewMongoStore(db)
		}
	} else {
		return store.NewMemoryStore()
	}
}

func configureEventQueues() {
	events.AddQueue(&events.CLIQueue{})
}

func configureAPIHandler(app *fiber.App, taskStore internal.TaskStore) fiber.Router {
	api := app.Group("/api")
	taskHandler := handlers.NewTaskHandler(taskStore)

	api.Get("/tasks", taskHandler.List)
	api.Get("/tasks/:id", taskHandler.Get)
	api.Post("/tasks", taskHandler.Create)
	api.Put("/tasks/:id", taskHandler.Update)
	api.Delete("/tasks/:id", taskHandler.Delete)

	return api
}
