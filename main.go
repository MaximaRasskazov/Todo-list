package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	portPtr := flag.Int("port", 3000, "port number")
	dirPtr := flag.String("dir", ".", "directory to serve")
	flag.Parse()

	// Формируем строку для прослушивания (например, ":3000")
	listenAddr := fmt.Sprintf(":%d", *portPtr)

	// Настраиваем FileServer для обслуживания статических файлов из указанной директории
	fs := http.FileServer(http.Dir(*dirPtr))
	http.Handle("/", fs)

	log.Printf("Сервер запущен на http://localhost:%d\n", *portPtr)
	log.Printf("Обслуживается директория: '%s'\n", *dirPtr)
	log.Println("Для остановки сервера нажмите Ctrl+C")

	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
