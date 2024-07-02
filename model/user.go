package model

// Модель для описания пользователя
type User struct {
	*UserFromJson
	ID         int    `json:"userID" db:"user_id"`
	Surname    string `json:"surname" db:"surname"`
	Name       string `json:"name" db:"name"`
	Patronymic string `json:"patronymic" db:"patronymic"`
	Address    string `json:"address" db:"address"`
}

// Метод для слияния нетронутых полей и измененных полей
func (u *User) MergeUpdates(userMap map[string]interface{}) {
	for key, value := range userMap {
		if key == "surname" && value.(string) != u.Surname {
			u.Surname = value.(string)
		} else if key == "name" && value.(string) != u.Name {
			u.Name = value.(string)
		} else if key == "patronymic" && value.(string) != u.Patronymic {
			u.Patronymic = value.(string)
		} else if key == "address" && value.(string) != u.Address {
			u.Address = value.(string)
		} else if key == "passportNumber" && value.(string) != u.PassportNumber {
			u.PassportNumber = value.(string)
		}
	}
}
