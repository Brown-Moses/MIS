package main

import (
	"MIS/database"
	"MIS/handler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	database.Connection()

	fmt.Println("Connection successful...")

	r := chi.NewRouter()

	r.Get("/student", handler.GetStudentsHandler)

	r.Get("/student/{id}", handler.GetStudentsByIDHandler)

	r.Post("/student", handler.StudentHandler)

	r.Post("/teacher", handler.TeacherHandler)

	http.ListenAndServe(":8080", r)
}
