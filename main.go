package main

import (
	"MIS/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	r.Post("/student", handler.StudentHandler)

	r.Post("/teacher", handler.TeacherHandler)

	http.ListenAndServe(":8080", r)
}
