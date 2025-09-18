package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Структуры данных
type Todo struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}

// Глобальные переменные
var todos []Todo
var currentID int

// Middleware
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next(w, r)
	}
}

// func logRequest(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Printf("%s  %s\n", r.Method, r.URL.Path)
// 		next(w, r)
// 	}
// }

// Обработчики
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	// Куда мы пишем? В w
	// Что кодируем? todos - глобальную переменную (в нее сохр. запросы)
	json.NewEncoder(w).Encode(todos)
}

func postTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	// Разгрузим наш запрос в newTodo
	var newTodo Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Вдруг нет обязательного поля
	if newTodo.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	currentID++
	newTodo.Id = currentID
	newTodo.CreatedAt = time.Now()
	newTodo.Completed = false // По умолч. задача не может быть выполнена

	todos = append(todos, newTodo)

	// Вернем теперь нашу собранную задачу
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

// Главная функция
func main() {
	// Инициализация данных
	todos = []Todo{
		{Id: 1, Title: "Привет", Completed: false, CreatedAt: time.Now()}, {Id: 2, Title: "Сосед", Completed: true, CreatedAt: time.Now()},
	}

	// !!! При запуске перезаписывается !!!
	currentID = 2

	portPtr := flag.Int("port", 3000, "номер порта")
	dirPtr := flag.String("dir", "./static", "директория на выгруз")
	flag.Parse()

	listenAddr := fmt.Sprintf(":%d", *portPtr)

	// Настраиваем обработчик для статических файлов
	fs := http.FileServer(http.Dir(*dirPtr))
	http.Handle("/", fs)

	// Регистрация маршрутов
	http.HandleFunc("/api/todos", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTodos(w, r)
		case http.MethodPost:
			postTodo(w, r)
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	log.Printf("Сервер запущен на http://localhost:%d\n", *portPtr)
	log.Printf("Обслуживается директория: '%s'\n", *dirPtr)
	log.Println("Для остановки сервера нажмите Ctrl+C")

	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
