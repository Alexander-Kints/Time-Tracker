package handler

import (
	"TimeTracker/model"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
)

// Handler для получения юзера по его id
func GetUserDyIDHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User

		id, err := strconv.Atoi(r.URL.Query().Get("userID"))
		if err != nil || id == 0 {
			JsonResponse(w, "userID is no valid", http.StatusBadRequest)
			return
		}

		query := `SELECT * FROM users WHERE user_id=$1`
		if err := db.Get(&user, query, id); err != nil {
			JsonResponse(w, "data error or user not exist", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		resp, err := json.Marshal(&user)
		if err != nil {
			JsonResponse(w, "server error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		if _, err := w.Write(resp); err != nil {
			JsonResponse(w, "server error", http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
}
