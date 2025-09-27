package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MaximaRasskazov/to-do-list/internal/database"
	"github.com/MaximaRasskazov/to-do-list/internal/handlers"
	"github.com/joho/godotenv"
)

// Middleware
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// Главная функция
func main() {
	// Загружаем переменные окружения
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Инициализируем базу данных
	if err := database.Init(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer database.Close()

	// Обрабатываем флаги командной строки
	portPtr := flag.Int("port", 3000, "номер порта")
	dirPtr := flag.String("dir", "./static", "директория для статических файлов")
	flag.Parse()

	// Используем порт из переменных окружения, если он задан
	// Иначе используем значение из флагов
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = fmt.Sprintf("%d", *portPtr)
	}

	listenAddr := ":" + port

	// Настраиваем обработчик для статических файлов
	fs := http.FileServer(http.Dir(*dirPtr))
	http.Handle("/", fs)

	// Регистрация маршрутов API
	http.HandleFunc("/api/todos", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetTodosHandler(w, r)
		case http.MethodPost:
			handlers.PostTodoHandler(w, r) // Обратите внимание на имя!
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	http.HandleFunc("/api/todos/", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			handlers.UpdateTodoHandler(w, r)
		case http.MethodDelete:
			handlers.DeleteTodoHandler(w, r)
		case http.MethodOptions:
			w.WriteHeader(http.StatusOK)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	log.Printf("Сервер запущен на http://localhost:%s\n", port)
	log.Printf("Обслуживается директория: '%s'\n", *dirPtr)
	log.Println("Для остановки сервера нажмите Ctrl+C")

	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
