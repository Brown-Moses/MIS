package main

import (
	"MIS/internal/database"
	"fmt"
	"net/http"

	"MIS/internal/handler"

	"log"

	"github.com/go-chi/cors"

	"github.com/go-chi/chi/v5"
)

func main() {
	//database connection
	database.Connect()

	fmt.Println("Connection successful...")

	//add a chi router
	r := chi.NewRouter()

	// inside main.go before routes
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	//student path group
	r.Route("/student", func(r chi.Router) {

		r.Get("/", handler.GetStudentHandler)

		r.Get("/{id}", handler.GetStudentByIDHandler)

		r.Post("/", handler.AddStudentHandler)

	})

	//teachers path group
	r.Route("/teacher", func(r chi.Router) {
		r.Post("/", handler.TeacherHandler)
	})

	log.Println("Server running on port 8080")

	//add error handling for things that go sideways
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("server failed %v", err)
	}
}
