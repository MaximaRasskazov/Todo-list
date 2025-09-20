package models

import "time"

// Todo структура задачи
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}

// Глобальные переменные (временное решение до подключения БД)
var (
	Todos     []Todo
	CurrentID int
)
