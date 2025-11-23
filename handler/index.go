package handler

import (
	student_handler "MIS/handler/student_handler"
	teacher_handler "MIS/handler/teacher_handler"
)

// student handlers
var (
	GetStudentHandler     = student_handler.GetStudentsHandler
	GetStudentByIDHandler = student_handler.GetStudentsByIDHandler
	StudentHandler        = student_handler.StudentHandler
)

// teachers
var (
	TeacherHandler = teacher_handler.TeachersHandler
)

//add more if there is!!!
