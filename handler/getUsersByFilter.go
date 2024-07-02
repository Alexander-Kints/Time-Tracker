package handler

import (
	"TimeTracker/model"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

// Handler для получения списка юзеров по какому-то фильтру.
// Работает, если поле json = поле db
func GetUsersByFilterHandler(db *sqlx.DB, filter string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userData := &struct {
			Users []*model.User `json:"users"`
		}{}

		filterParam := r.URL.Query().Get(filter)

		if filterParam == "" {
			JsonResponse(w, fmt.Sprintf("%s is no valid", filter), http.StatusBadRequest)
			return
		}

		query := fmt.Sprintf(`SELECT * FROM users WHERE %s=$1`, filter)
		if err := db.Select(&userData.Users, query, filterParam); err != nil {
			JsonResponse(w, "data error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

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

		log.Printf("send %d users by filter %s and param %s", len(userData.Users), filter, filterParam)
	}
}
