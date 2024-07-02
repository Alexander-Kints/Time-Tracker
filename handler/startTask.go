package handler

import (
	"TimeTracker/model"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"
)

// Handler для создания задачи
func StartTaskHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task model.Task

		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			JsonResponse(w, "server error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		if err := task.CheckData(); err != nil {
			JsonResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		task.StartedAt = time.Now()

		query := `INSERT INTO tasks (is_completed, title, user_id, started_at)
					VALUES($1, $2, $3, $4) RETURNING task_id`

		row := db.QueryRow(query, task.IsCompleted, task.Title, task.UserID, task.StartedAt)

		var id int
		err := row.Scan(&id)
		if err != nil {
			JsonResponse(w, fmt.Sprintf("user id %d not exist", task.UserID), http.StatusBadRequest)
			log.Println(err)
			return
		}

		JsonResponse(w, fmt.Sprintf("task id %d created", id), http.StatusOK)
		log.Printf("task id %d created", id)
	}
}
