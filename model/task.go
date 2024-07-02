package model

import (
	"errors"
	"time"
)

// Модель для описания задачи
type Task struct {
	ID          int       `json:"taskID" db:"task_id"`
	IsCompleted bool      `json:"isCompleted" db:"is_completed"`
	Title       string    `json:"title" db:"title"`
	UserID      int       `json:"userID" db:"user_id"`
	StartedAt   time.Time `json:"startAt" db:"started_at"`
	FinishedAt  time.Time `json:"finishedAt" db:"finished_at"`
	Duration    string    `json:"duration" db:"duration"` // затраченное время в часах и минутах
}

// Метод для валидации полей, используется при создании задачи
func (t *Task) CheckData() error {
	if t.UserID == 0 {
		return errors.New("empty userID")
	}

	if t.Title == "" {
		return errors.New("empty title")
	}

	return nil
}
