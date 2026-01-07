package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Структура модели
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main2() {
	mux := http.NewServeMux()

	// Ручка GET /users/{id}, возвращает JSON объект пользователя
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		userIDStr := r.URL.Path[len("/users/"):]
		if userIDStr == "" || len(userIDStr) > 8 { // Простая проверка валидности ID
			http.Error(w, "Invalid user id", http.StatusBadRequest)
			return
		}

		// Предположим, мы знаем заранее пользователя с указанным ID
		var user User
		switch userIDStr {
		case "1":
			user = User{ID: 1, Name: "Иван"}
		default:
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&user)
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
