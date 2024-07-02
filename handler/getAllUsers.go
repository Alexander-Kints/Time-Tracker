package handler

import (
	"TimeTracker/model"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
)

// Handler для получения списка юзеров с пагинация (с {число} по {число})
func GetAllUsersHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userData := &struct {
			From  int           `json:"from"`
			To    int           `json:"to"`
			Users []*model.User `json:"users"`
		}{}

		from, err := strconv.Atoi(r.URL.Query().Get("from"))
		if err != nil {
			JsonResponse(w, "from is no valid", http.StatusBadRequest)
			return
		}
		to, err := strconv.Atoi(r.URL.Query().Get("to"))
		if err != nil {
			JsonResponse(w, "to is no valid", http.StatusBadRequest)
			return
		}

		limit := (to - from) + 1

		query := `SELECT * FROM users LIMIT $1 OFFSET $2`
		if err := db.Select(&userData.Users, query, limit, from-1); err != nil {
			JsonResponse(w, "data error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		userData.From = from
		userData.To = to

		resp, err := json.Marshal(userData)
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
		log.Printf("send users from %d to %d", from, to)
	}
}
