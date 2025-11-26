package main

import (
	"MIS/internal/database"
	"fmt"
	"net/http"

	"MIS/internal/handler"

	"log"

	"github.com/go-chi/chi/v5"
)

func main() {
	//database connection
	database.Connect()

	fmt.Println("Connection successful...")

	//add a chi router
	r := chi.NewRouter()

	//student path group
	r.Route("/student", func(r chi.Router) {

		r.Get("/student", handler.GetStudentHandler)

		r.Get("/student/{id}", handler.GetStudentByIDHandler)

		r.Post("/student", handler.StudentHandler)

	})

	//teachers path group
	r.Route("/teacher", func(r chi.Router) {
		r.Post("/teacher", handler.TeacherHandler)
	})

	log.Println("Server running on port 8080")

	//add error handling for things that go sideways
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("server failed %v", err)
	}
}
