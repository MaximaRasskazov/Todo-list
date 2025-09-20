package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// GetTodosHandler обработчик для получения всех задач
func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// CreateTodoHandler обработчик для создания новой задачи
func PostTodoHandler(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	currentID++
	newTodo.ID = currentID
	newTodo.CreatedAt = time.Now()
	todos = append(todos, newTodo)

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
	var todo *Todo
	for i := range todos {
		if todos[i].ID == id {
			todo = &todos[i]
			break
		}
	}

	if todo == nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Обновляем задачу
	var updatedTodo Todo
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
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}
