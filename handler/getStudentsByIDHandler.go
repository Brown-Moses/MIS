package handler

import (
	"MIS/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetStudentsByIDHandler(w http.ResponseWriter, r *http.Request) {

	//set headers
	w.Header().Set("Content-Type", "application/json")

	//validates the get method
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract student ID from URL (e.g., /students/123)
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	// Convert string to uint
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}

	// Pass the student ID to the service
	studentsById, err := service.GetStudentsByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// encode from go to json
	json.NewEncoder(w).Encode(studentsById)

}
