package main

import (
	"os"

	"github.com/bespinian/ict-todo/backend/tasks/internal"
	"github.com/bespinian/ict-todo/backend/tasks/internal/events"
	"github.com/bespinian/ict-todo/backend/tasks/internal/handlers"
	"github.com/bespinian/ict-todo/backend/tasks/internal/store"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/redis/go-redis/v9"
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
		log.Info("MONGODB_URI is set. Using MongoDB as storage.")
		db, err := store.NewMongoDatabase(mongoUri, "ict")
		if err != nil {
			log.Infof("Could not connect to MongoDB: %s. Falling back to in-memory store.", err)
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

	if url := os.Getenv("REDIS_URL"); url != "" {
		log.Info("REDIS_URL is set. Using Redis event queue.")
		opts, err := redis.ParseURL(url)
		if err != nil {
			log.Error("Could not parse REDIS_URL.")
			opts = &redis.Options{}
		}
		redisClient := redis.NewClient(opts)
		redisQueue := events.NewRedisQueue(redisClient)
		events.AddQueue(redisQueue)
	}
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
