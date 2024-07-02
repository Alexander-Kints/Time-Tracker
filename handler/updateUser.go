package handler

import (
	"TimeTracker/model"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

// Handler для обновления данных юзера по его id
func UpdateUserHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		var userMap map[string]interface{}

		// Декодирование json в map
		if err := json.NewDecoder(r.Body).Decode(&userMap); err != nil {
			JsonResponse(w, "server error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		// Если в map нет id, отправка ошибки
		id, ok := userMap["userID"]
		if !ok {
			JsonResponse(w, "userID is no valid", http.StatusBadRequest)
			return
		}

		// Получения текущих данных пользователя
		query := `SELECT * FROM users WHERE user_id=$1`
		if err := db.Get(&user, query, id); err != nil {
			JsonResponse(w, "data error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		// Слияние нетронутых полей и измененных полей
		user.MergeUpdates(userMap)

		query = `UPDATE users SET passport_number=$1, surname=$2, name=$3, patronymic=$4, address=$5 WHERE user_id=$6`
		if _, err := db.Exec(query, user.PassportNumber, user.Surname, user.Name, user.Patronymic, user.Address, user.ID); err != nil {
			JsonResponse(w, "data error or user not exist", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		JsonResponse(w, fmt.Sprintf("user id %d was updated", user.ID), http.StatusBadRequest)
		log.Printf("user id %d was updated", user.ID)
	}
}
