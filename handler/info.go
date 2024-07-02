package handler

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
)

// эндпоинт, описанный в сваггере (имитация обращения к стороннему сервису)
func InfoHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		passportSeries, err := strconv.Atoi(r.URL.Query().Get("passportSerie"))
		if err != nil {
			JsonResponse(w, "passportSerie is no valid", http.StatusBadRequest)
			return
		}

		passportNumber, err := strconv.Atoi(r.URL.Query().Get("passportNumber"))
		if err != nil {
			JsonResponse(w, "passportNumber is no valid", http.StatusBadRequest)
			return
		}

		info := &struct {
			Surname    string `json:"surname" db:"surname"`
			Name       string `json:"name" db:"name"`
			Patronymic string `json:"patronymic" db:"patronymic"`
			Address    string `json:"address" db:"address"`
		}{}

		query := `SELECT surname, name, patronymic, address
				FROM info
				WHERE passport_series=$1 AND passport_number=$2`

		if err := db.Get(info, query, passportSeries, passportNumber); err != nil {
			JsonResponse(w, "data error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		resp, err := json.Marshal(&info)
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
