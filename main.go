package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

type MeResponse struct {
	Name	string 	`json:"name"`
	Email	string	`json:"email"`
	Github	string 	`json:"github"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Message: "API is running"})
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Message: "healthy"})
	})

	http.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(MeResponse{
			Name: "Tiza Talla",
			Email: "tallatiza6@gmail.com",
			Github: "https://github.com/kurocifer",
		})
	})

	http.ListenAndServe(":8080", nil)
}