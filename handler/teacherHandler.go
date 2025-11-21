package handler

import (
	"MIS/model"
	"MIS/service"
	"encoding/json"
	"net/http"
)

func TeacherHandler(w http.ResponseWriter, r *http.Request) {

	var t model.Teacher

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
