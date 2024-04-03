package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Встановлення маршрутів для відслідковування запитів
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/City3d.json", serveJSON)

	// Запуск веб-сервера на порту 8000
	fmt.Println("Server is running on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

// Функція для обробки запитів до головної сторінки
func serveIndex(w http.ResponseWriter, r *http.Request) {
	// Відправляємо відповідь клієнту
	http.ServeFile(w, r, "newIndex2.html")
}

// Функція для обробки запитів до файлу City.json
func serveJSON(w http.ResponseWriter, r *http.Request) {
	// Відправляємо відповідь клієнту
	http.ServeFile(w, r, "City3d.json")
}
