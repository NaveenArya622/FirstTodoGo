package dbHelper

import (
	"database/sql"
	"fmt"
	"time"
	"todo/database"
	"todo/models"
)

func CreateTask(userId, name, description string) (string, error) {
	// language=SQL
	SQL := `INSERT INTO todo(user_id, name, description) VALUES ($1, $2, $3) RETURNING id`
	var taskID string
	if err := database.Todo.QueryRowx(SQL, userId, name, description).Scan(&taskID); err != nil {
		return "", err
	}
	return taskID, nil
}

func TaskCompleted(userId, taskId string) (string, error) {
	// language=SQL
	SQL := `UPDATE 
				todo t 
			SET 
				pending_at= $1 
			WHERE 
				t.pending_at IS NULL AND 
				t.archived_at IS NULL AND 
				t.user_id = $2 
				AND t.id = $3`
	_, err := database.Todo.Exec(SQL, time.Now(), userId, taskId)
	if err != nil {
		return "", err
	}
	return taskId, nil
}

func TaskArchived(userId, taskId string) (string, error) {
	// language=SQL
	SQL := `UPDATE 
				todo t 
			SET 
				archived_at= $1 
			WHERE 
				t.archived_at IS NULL AND 
				t.user_id = $2 
				AND t.id = $3`
	_, err := database.Todo.Exec(SQL, time.Now(), userId, taskId)
	if err != nil {
		return "", err
	}
	return taskId, nil
}

func GetTasksByUserID(userId string) ([]models.Task, error) {
	// language=SQL
	SQL := `SELECT
				t.id,
				t.user_id,
				t.name,
				t.description,
				t.created_at,
				t.pending_at
       		FROM
				todo t
			WHERE 
				t.archived_at IS NULL AND
				t.user_id=$1`
	tasks := make([]models.Task, 0)
	err := database.Todo.Select(&tasks, SQL, userId)
	fmt.Println(tasks, userId)
	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return tasks, nil
}

func GetTasksByID(userId, taskId string) (*models.Task, error) {
	// language=SQL
	SQL := `SELECT
				t.id,
				t.user_id,
				t.name,
				t.description,
				t.created_at,
				t.pending_at
       		FROM
				todo t
			WHERE t.archived_at IS NULL AND t.user_id=$1 AND t.id=$2`
	var taskById models.Task
	err := database.Todo.Get(&taskById, SQL, userId, taskId)
	if err != nil {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &taskById, nil
}
