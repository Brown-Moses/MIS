package handler

import (
	student_handler "MIS/internal/handler/student_handler"
	teacher_handler "MIS/internal/handler/teacher_handler"
)

// student handlers
var (
	GetStudentHandler     = student_handler.GetStudentsHandler
	GetStudentByIDHandler = student_handler.GetStudentsByIDHandler
	AddStudentHandler     = student_handler.AddStudentHandler
)

// teachers
var (
	TeacherHandler = teacher_handler.TeachersHandler
)

//add more if there is!!!
