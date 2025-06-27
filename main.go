package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(Response{Message: "Hello World"})
		if err != nil {
			panic(err)
		}
	})

	url := "localhost:3000"
	fmt.Println("Server started on", url)
	err := http.ListenAndServe(url, r)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

}
