package handler

import (
	"fmt"
	"net/http"
	"todo/database/dbHelper"
	"todo/middlewares"
	"todo/models"
	"todo/utils"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}{}
	if parseErr := utils.ParseBody(r.Body, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}
	fmt.Println(body, middlewares.UserContext(r).ID)

	taskID, saveErr := dbHelper.CreateTask(middlewares.UserContext(r).ID, body.Name, body.Description)
	if saveErr != nil {
		return
	}
	utils.RespondJSON(w, http.StatusCreated, struct {
		TaskID string `json:"taskID"`
	}{
		TaskID: taskID,
	})
}

func GetTasksByUserID(w http.ResponseWriter, r *http.Request) {
	tasks, saveErr := dbHelper.GetTasksByUserID(middlewares.UserContext(r).ID)
	if saveErr != nil {
		return
	}
	utils.RespondJSON(w, http.StatusCreated, struct {
		Tasks []models.Task `json:"tasks"`
	}{
		Tasks: tasks,
	})
}

func GetTasksByID(w http.ResponseWriter, r *http.Request) {
	body := struct {
		ID string `json:"id"`
	}{}
	if parseErr := utils.ParseBody(r.Body, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}

	task, saveErr := dbHelper.GetTasksByID(middlewares.UserContext(r).ID, body.ID)
	if saveErr != nil {
		return
	}
	utils.RespondJSON(w, http.StatusCreated, struct {
		Task models.Task `json:"task"`
	}{
		Task: *task,
	})
}

func TaskCompleted(w http.ResponseWriter, r *http.Request) {
	body := struct {
		ID string `json:"id"`
	}{}
	if parseErr := utils.ParseBody(r.Body, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}

	taskId, saveErr := dbHelper.TaskCompleted(middlewares.UserContext(r).ID, body.ID)
	if saveErr != nil {
		return
	}
	utils.RespondJSON(w, http.StatusCreated, struct {
		TaskID string `json:"taskID"`
	}{
		TaskID: taskId,
	})
}

func TaskArchived(w http.ResponseWriter, r *http.Request) {
	body := struct {
		ID string `json:"id"`
	}{}
	if parseErr := utils.ParseBody(r.Body, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}

	taskId, saveErr := dbHelper.TaskArchived(middlewares.UserContext(r).ID, body.ID)
	if saveErr != nil {
		return
	}
	utils.RespondJSON(w, http.StatusCreated, struct {
		TaskID string `json:"taskID"`
	}{
		TaskID: taskId,
	})
}
