package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var messages []string

func getMessages(w http.ResponseWriter, r *http.Request) {
	for _, message := range messages {
		fmt.Fprintln(w, message)
	}
}

func addMessage(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue("message")
	messages = append(messages, msg)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/messages", getMessages).Methods("GET")
	router.HandleFunc("/messages", addMessage).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
