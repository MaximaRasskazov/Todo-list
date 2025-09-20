package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/MaximaRasskazov/to-do-list/internal/models"
)

// GetTodosHandler обработчик для получения всех задач
func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Todos)
}

// CreateTodoHandler обработчик для создания новой задачи
func PostTodoHandler(w http.ResponseWriter, r *http.Request) {
	var newTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	models.CurrentID++
	newTodo.ID = models.CurrentID
	newTodo.CreatedAt = time.Now()
	models.Todos = append(models.Todos, newTodo)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTodo)
}

// UpdateTodoHandler - обновление задачи
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID из URL: /api/todos/42 -- получаем 42
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/todos/"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Находим задачу
	var todo *models.Todo
	for i := range models.Todos {
		if models.Todos[i].ID == id {
			todo = &models.Todos[i]
			break
		}
	}

	if todo == nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Обновляем задачу
	var updatedTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	todo.Title = updatedTodo.Title
	todo.Completed = updatedTodo.Completed

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// DeleteTodoHandler обработчик для удаления задачи
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/todos/"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Находим и удаляем задачу
	for i, todo := range models.Todos {
		if todo.ID == id {
			models.Todos = append(models.Todos[:i], models.Todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}
