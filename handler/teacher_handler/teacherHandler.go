package handler

import (
	"MIS/model"
	service "MIS/service/teacher_service"
	"encoding/json"
	"net/http"
)

func TeachersHandler(w http.ResponseWriter, r *http.Request) {

	var t model.Teacher

	//valid post method
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//decode the json to go
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	//ensure the validation too
	savedTeacher, err := service.AddTeacher(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//respond
	//start with header
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(savedTeacher)

}
