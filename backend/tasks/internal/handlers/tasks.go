package handlers

import (
	"github.com/bespinian/ict-todo/backend/tasks/internal"
	"github.com/bespinian/ict-todo/backend/tasks/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type TaskHandler struct {
	store internal.TaskStore
}

func NewTaskHandler(store internal.TaskStore) *TaskHandler {
	return &TaskHandler{store}
}

func (h *TaskHandler) List(c *fiber.Ctx) error {
	tasks := h.store.List()
	return c.JSON(tasks)
}

func (h *TaskHandler) Get(c *fiber.Ctx) error {
	taskId := c.Params("id")

	task := h.store.Get(taskId)
	if task == nil {
		return fiber.ErrNotFound
	}
	return c.JSON(task)
}

func (h *TaskHandler) Create(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		log.Error(err)
		return err
	}

	newTask, err := h.store.Add(task)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(newTask)
}

func (h *TaskHandler) Update(c *fiber.Ctx) error {
	return c.JSON(models.Task{})
}

func (h *TaskHandler) Delete(c *fiber.Ctx) error {
	return c.SendStatus(204)
}
