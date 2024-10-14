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
	tasks := h.store.List(c.Context())
	return c.JSON(tasks)
}

func (h *TaskHandler) Get(c *fiber.Ctx) error {
	taskId := c.Params("id")

	task := h.store.Get(c.Context(), taskId)
	if task == nil {
		return fiber.ErrNotFound
	}
	return c.JSON(task)
}

func (h *TaskHandler) Create(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return err
	}

	newTask, err := h.store.Add(c.Context(), task)
	if err != nil {
		log.Error(err)
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(newTask)
}

func (h *TaskHandler) Update(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return err
	}

	if err := h.store.Update(c.Context(), task); err != nil {
		log.Error(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(task)
}

func (h *TaskHandler) Delete(c *fiber.Ctx) error {
	if err := h.store.Delete(c.Context(), c.Params("id")); err != nil {
		return err
	}
	return c.SendStatus(204)
}
