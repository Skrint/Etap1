package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string
var msg string

type requestBody struct {
	Message string `json:"message"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	message := requestBody{Message: "message"}
	tasks := requestBody{Message: "Hello, world!"}
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "Получено: %s и %s\n", message, tasks)
		task = string(tasks.Message)
		msg = string(message.Message)
	} else {
		fmt.Fprintln(w, "Поддерживается только метод POST")
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, `{"%s":"%s"}`, msg, task)
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
