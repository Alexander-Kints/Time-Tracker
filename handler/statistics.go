package handler

import (
	"TimeTracker/model"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
)

// Handler для получения статистики юзера по его id
func StatisticsHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		var statistics model.Statistics

		userID, err := strconv.Atoi(r.URL.Query().Get("userID"))
		if err != nil || userID == 0 {
			JsonResponse(w, "userID is no valid", http.StatusBadRequest)
			return
		}

		query := `SELECT surname, name, patronymic FROM users WHERE user_id=$1`
		if err := db.Get(&user, query, userID); err != nil {
			JsonResponse(w, "data error or user not exist", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		query = `SELECT * FROM tasks WHERE user_id=$1 AND is_completed ORDER BY duration DESC`

		err = db.Select(&(statistics.Tasks), query, userID)
		if err != nil {
			JsonResponse(w, "data error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		statistics.UserID = userID
		statistics.FullName = fmt.Sprintf("%s %s %s", user.Surname, user.Name, user.Patronymic)
		statistics.CompletedTasks = len(statistics.Tasks)

		resp, err := json.Marshal(&statistics)
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

		log.Printf("send statistics user id %d", userID)
	}
}
