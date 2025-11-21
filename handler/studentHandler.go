package handler

import (
	"MIS/model"
	"MIS/service"
	"encoding/json"
	"net/http"
)

func StudentHandler(w http.ResponseWriter, r *http.Request) {

	var s model.Student

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
