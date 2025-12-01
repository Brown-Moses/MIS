package handler

import (
	"MIS/internal/dto"
	service "MIS/internal/service/student_service"
	"encoding/json"
	"net/http"
)

func AddStudentHandler(w http.ResponseWriter, r *http.Request) {

	// Parse form values from HTMX
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}

	s := dto.StudentDTO{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Gender:    r.FormValue("gender"),
		DOB:       r.FormValue("date_of_birth"),
	}

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
