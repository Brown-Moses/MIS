package handler

import (
	"MIS/model"
	service "MIS/service/student_service"
	"encoding/json"
	"net/http"
)

func StudentHandler(w http.ResponseWriter, r *http.Request) {

	var s model.Student

	//valid post method
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//decoding next and save to s
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	//add service
	savedStudent, err := service.AddStudent(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//respond, we encode then from go to json
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(savedStudent)

}
