package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/MaximaRasskazov/to-do-list/internal/database"
	"github.com/MaximaRasskazov/to-do-list/internal/models"
)

// Middleware
func EnableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next(w, r)
	}
}

// GetTodosHandler обработчик для получения всех задач
func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(r.Context(),
		"SELECT id, title, completed, created_at, updated_at FROM todos ORDER BY created_at DESC")
	if err != nil {
		http.Error(w, "Error fetching todos: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			http.Error(w, "Error scanning todo: "+err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// CreateTodoHandler обработчик для создания новой задачи
func PostTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if todo.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	// Вставляем задачу в БД и получаем обратно сгенерированные ID и даты
	err := database.DB.QueryRow(r.Context(),
		"INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id, created_at, updated_at",
		todo.Title, todo.Completed).Scan(&todo.ID, &todo.CreatedAt, &todo.UpdatedAt)

	if err != nil {
		http.Error(w, "Error creating todo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

// UpdateTodoHandler - обновление задачи
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID из URL: /api/todos/42 -- получаем 42
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/todos/"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Декодируем JSON из тела запроса
	var updatedTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Обновляем задачу в базе данных
	result, err := database.DB.Exec(r.Context(),
		"UPDATE todos SET title = $1, completed = $2, updated_at = NOW() WHERE id = $3",
		updatedTodo.Title, updatedTodo.Completed, id)

	if err != nil {
		http.Error(w, "Error updating todo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Проверяем, была ли обновлена хотя бы одна строка
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Получаем обновлённую задачу для возврата в ответе
	var todo models.Todo
	err = database.DB.QueryRow(r.Context(),
		"SELECT id, title, completed, created_at, updated_at FROM todos WHERE id = $1",
		id).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)

	if err != nil {
		http.Error(w, "Error fetching updated todo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// DeleteTodoHandler обработчик для удаления задачи
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID из URL: /api/todos/42 -- получаем 42
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/todos/"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Выполняем удаление задачи из базы данных
	result, err := database.DB.Exec(r.Context(),
		"DELETE FROM todos WHERE id = $1", id)

	if err != nil {
		http.Error(w, "Error deleting todo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Проверяем, была ли удалена хотя бы одна строка
	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Возвращаем статус 204 No Content при успешном удалении
	w.WriteHeader(http.StatusNoContent)
}
