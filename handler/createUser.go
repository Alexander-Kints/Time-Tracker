package handler

import (
	"TimeTracker/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strings"
)

// Handler для создания юзера
func CreateUserHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userFromJson model.UserFromJson
		var id int

		// Парсинг json-объекта, возврат json с сообщением ошибки в случае ошибки
		if err := json.NewDecoder(r.Body).Decode(&userFromJson); err != nil {
			JsonResponse(w, "server error", http.StatusBadRequest)
			log.Println(err)
			return
		}

		// Валидация json-объекта, возврат json с сообщением ошибки в случае ошибки
		if err := userFromJson.CheckData(); err != nil {
			JsonResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Обращение к сервису для получения данных юзера по его паспорту
		// Можно сделать через горутину и контекст
		user, err := infoRequest("http://127.0.0.1:9000/info", &userFromJson)
		if err != nil {
			JsonResponse(w, "server error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		query := `INSERT INTO users 
    		(passport_number, surname, name, patronymic, address)
			VALUES ($1, $2, $3, $4, $5) RETURNING user_id`
		row := db.QueryRow(query, user.PassportNumber, user.Surname, user.Name, user.Patronymic, user.Address)
		if err := row.Scan(&id); err != nil {
			JsonResponse(w, "data error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		// Отправка сообщения об успехе операции
		JsonResponse(w, fmt.Sprintf("user id %d created", id), http.StatusOK)
		log.Printf("user id %d created", id)
	}
}

// Запрос к /info для получения данных юзера
func infoRequest(url string, userFromJson *model.UserFromJson) (*model.User, error) {
	client := &http.Client{}
	user := model.User{
		UserFromJson: userFromJson,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	query := req.URL.Query()
	params := strings.Split(userFromJson.PassportNumber, " ")
	query.Add("passportSerie", params[0])
	query.Add("passportNumber", params[1])
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("bad Request or user info not found in database")
	}

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
