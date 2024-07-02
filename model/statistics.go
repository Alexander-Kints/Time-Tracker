package model

// Модель для описания статистики юзера
type Statistics struct {
	UserID         int     `json:"userID"`
	FullName       string  `json:"fullName"`
	CompletedTasks int     `json:"completedTasks"` // кол-во выполненных задач
	Tasks          []*Task `json:"tasks"`
}
