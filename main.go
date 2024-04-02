package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Граф
var graph = map[string]map[string]int{
	"A": {"B": 10, "C": 15, "E": 5},
	"B": {"D": 12, "C": 10},
	"C": {"D": 5},
	"D": {"E": 7},
	"E": {"D": 10},
}

// Обробник запитів
func handleDistanceRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Парсинг запиту
	var requestData struct {
		Start string `json:"start"`
		End   string `json:"end"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Пошук відстані між містами
	distance, ok := graph[requestData.Start][requestData.End]
	if !ok {
		distance = -1 // Якщо відстань не знайдено
	}

	// Відповідь
	responseData := map[string]int{"distance": distance}
	json.NewEncoder(w).Encode(responseData)
}

func main() {
	// Обробник запитів
	http.HandleFunc("/distance", handleDistanceRequest)

	// Запуск веб-сервера
	fmt.Println("Сервер запущений на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
