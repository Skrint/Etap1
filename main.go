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
	if r.Method == http.MethodPost {
		var reqBody requestBody
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, "Ошибка декодирования:", http.StatusBadRequest)
			return
		}
		task = reqBody.Message
		fmt.Fprintf(w, "Получено: %s\n", task)
	} else {
		fmt.Fprintln(w, "Поддерживается только метод POST")
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Hello, %s\n", task)
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
