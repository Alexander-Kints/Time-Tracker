package handler

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
)

// Handler для удаления юзера по его id
func DeleteUserHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("userID"))
		if err != nil || id == 0 {
			JsonResponse(w, "userID is no valid", http.StatusBadRequest)
			return
		}

		query := `DELETE FROM users WHERE user_id=$1`
		_, err = db.Exec(query, id)
		if err != nil {
			JsonResponse(w, "data error or user not exist", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		JsonResponse(w, fmt.Sprintf("user id %d was deleted", id), http.StatusOK)
		log.Printf("user id %d was deleted", id)
	}
}
