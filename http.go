package main

import (
	"encoding/json"
	"net/http"
)

func startHTTP() {
	http.HandleFunc("/push", handlePush)
}

func handlePush(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var body pushRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = push(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
