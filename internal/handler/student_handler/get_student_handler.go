package handler

import (
	service "MIS/internal/service/student_service"
	"encoding/json"
	"net/http"
)

func GetStudentsHandler(w http.ResponseWriter, r *http.Request) {

	//validate method
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//get service
	getStudent, err := service.GetStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//set headers
	w.Header().Set("Content-Type", "Application/json")
	//encode
	json.NewEncoder(w).Encode(getStudent)
}
