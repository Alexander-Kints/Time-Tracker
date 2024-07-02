package model

import (
	"errors"
	"strconv"
	"strings"
)

// Модель для описания данных юзера, которые приходят при его создании
type UserFromJson struct {
	PassportNumber string `json:"passportNumber" db:"passport_number"`
}

// Метод для валидации полей, используется при создании юзера
func (u *UserFromJson) CheckData() error {
	if u.PassportNumber == "" {
		return errors.New("passportNumber is empty")
	}
	if args := strings.Split(u.PassportNumber, " "); len(args[0]) != 4 || len(args[1]) != 6 {
		return errors.New("passportNumber is no valid")
	} else if _, err := strconv.Atoi(args[0]); err != nil {
		return errors.New("series is no valid")
	} else if _, err := strconv.Atoi(args[1]); err != nil {
		return errors.New("number is no valid")
	}

	return nil
}
