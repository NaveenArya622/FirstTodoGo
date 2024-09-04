package server

import (
	"todo/handler"

	"github.com/go-chi/chi/v5"
)

func taskRoutes(r chi.Router) {
	r.Group(func(task chi.Router) {
		task.Post("/create", handler.CreateTask)
		task.Get("/details", handler.GetTasksByUserID)
		task.Get("/info", handler.GetTasksByID)
		task.Put("/", handler.TaskCompleted)
		task.Delete("/", handler.TaskArchived)
	})
}
