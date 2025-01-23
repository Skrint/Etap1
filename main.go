package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type requestBody struct {
	Message string `json:"message"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	taskRequestBody := requestBody{Message: "Hello, world!"}
	taskToJson, errTaskToJson := json.Marshal(taskRequestBody)
	if errTaskToJson != nil {
		fmt.Println("Ошибка маршалинга", errTaskToJson)
		return
	}
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "Получено: %s\n", taskToJson)
		task = string(taskToJson)
	} else {
		fmt.Fprintln(w, "Поддерживается только метод POST")
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, `%s`, task)
	} else {
		fmt.Fprintln(w, "Поддерживается только метод GET")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", GetHandler).Methods("GET")
	router.HandleFunc("/api/hello", PostHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
