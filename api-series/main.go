package main

import (
	"net/http"

	"api-series/db"
	"api-series/handlers"
)

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next(w, r)
	}
}

func main() {
	db.Connect()

	http.HandleFunc("/series", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			handlers.GetSeries(w, r)
		case "POST":
			handlers.CreateSeries(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	}))

	http.HandleFunc("/series/", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			handlers.GetSeriesByID(w, r)
		case "PUT":
			handlers.UpdateSeries(w, r)
		case "DELETE":
			handlers.DeleteSeries(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	}))

	http.ListenAndServe(":8080", nil)
}
