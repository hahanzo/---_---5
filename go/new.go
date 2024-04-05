package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

type City struct {
	Name        string       `json:"name"`
	Connections []Connection `json:"connections"`
}

type Connection struct {
	City     string `json:"city"`
	Distance int    `json:"distance"`
}

func main() {
	// Завантажуємо дані з JSON файлу
	data, err := os.ReadFile("C:\\Project\\СШІ_Лаб5\\SHI_lab5\\json\\distances.json")
	if err != nil {
		fmt.Println("Помилка читання файлу:", err)
		return
	}

	// Розпарсюємо дані JSON у структуру графа
	var city []City
	if err := json.Unmarshal(data, &city); err != nil {
		fmt.Println("Помилка парсингу JSON:", err)
		return
	}

	var graph = make(map[string][]Connection)
	for _, city := range city {
		graph[city.Name] = city.Connections
	}

	// Початкове та кінцеве місто
	startCity := "Вінниця"
	endCity := "Київ"

	// Пошук маршруту по жадібному алгоритму
	path, pathDistance := findGreedyPath(graph, startCity, endCity)

	// Виведення маршруту та загальної відстані
	fmt.Println("Маршрут:")
	for i := 0; i < len(path)-1; i++ {
		currentCity := path[i]
		nextCity := path[i+1]
		fmt.Printf("%s -> %s \n", currentCity, nextCity)
	}
	fmt.Println("Загальна дистанція:", pathDistance, "км")
}

// Пошук маршруту по жадібному алгоритму
func findGreedyPath(graph map[string][]Connection, startCity, endCity string) ([]string, int) {
	visited := make(map[string]bool)
	visited[startCity] = true
	path := []string{startCity}
	pathDistance := 0

	currentCity := startCity

	for currentCity != endCity {
		nearestCity := ""
		minDistance := math.MaxInt32

		// Пошук найближчого міста
		for _, connection := range graph[currentCity] {
			if !visited[connection.City] && connection.Distance < minDistance {
				nearestCity = connection.City
				minDistance = connection.Distance
			}
		}

		// Якщо найближче місто не знайдено, виходимо з циклу
		if nearestCity == "" {
			break
		}

		// Додаємо найближче місто до маршруту та позначаємо його як відвідане
		path = append(path, nearestCity)
		pathDistance += minDistance
		visited[nearestCity] = true
		currentCity = nearestCity
	}

	return path, pathDistance
}
