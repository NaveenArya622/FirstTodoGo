package models

import (
	"github.com/volatiletech/null"
)

type Task struct {
	ID          string    `json:"id" db:"id"`
	UserID      string    `json:"userId" db:"user_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   null.Time `json:"createdAt" db:"created_at"`
	PendingAt   null.Time `json:"pendingAt" db:"pending_at"`
}
