package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("welcome"))
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
