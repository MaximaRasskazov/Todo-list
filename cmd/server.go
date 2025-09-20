package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
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

// Главная функция
func main() {
	// Инициализация данных
	todos = []Todo{
		{ID: 1, Title: "Привет", Completed: false, CreatedAt: time.Now()}, {ID: 2, Title: "Сосед", Completed: true, CreatedAt: time.Now()},
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
			GetTodosHandler(w, r)
		case http.MethodPost:
			PostTodoHandler(w, r)
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	http.HandleFunc("/api/todos/", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			UpdateTodoHandler(w, r)
		case http.MethodDelete:
			DeleteTodoHandler(w, r)
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
