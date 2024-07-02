package handler

import (
	"TimeTracker/model"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Handler для завершения задачи по ее id
func FinishTaskHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task model.Task

		id, err := strconv.Atoi(r.URL.Query().Get("taskID"))
		if err != nil || id == 0 {
			JsonResponse(w, "taskID is no valid", http.StatusBadRequest)
			return
		}

		task.ID = id

		query := `SELECT is_completed, title, user_id, started_at FROM tasks WHERE task_id=$1`
		if err := db.Get(&task, query, task.ID); err != nil {
			JsonResponse(w, "data error or task not exist", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		// Если задача уже выполнена, возврат сообщения об этом и return
		if task.IsCompleted {
			JsonResponse(w, "task already completed", http.StatusBadRequest)
			return
		}

		task.FinishedAt = time.Now()
		task.Duration = task.FinishedAt.Sub(task.StartedAt).String()
		task.IsCompleted = true

		query = `UPDATE tasks SET is_completed=$1, finished_at=$2, duration=$3 WHERE task_id=$4`
		_, err = db.Exec(query, task.IsCompleted, task.FinishedAt, task.Duration, task.ID)
		if err != nil {
			JsonResponse(w, "data error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		resp, err := json.Marshal(&task)
		if err != nil {
			JsonResponse(w, "server error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		_, err = w.Write(resp)
		if err != nil {
			JsonResponse(w, "server error", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		log.Printf("task id %d was completed", task.ID)
	}
}
