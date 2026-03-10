package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(teacherHandler *TeacherHandler, studentHandler *StudentHandler, subjectHandler *SubjectHandler) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		// Teacher routes
		teachers := v1.Group("/teachers")
		{
			teachers.GET("", teacherHandler.GetAllTeachers)
			teachers.GET("/:id", teacherHandler.GetTeacher)
			teachers.POST("", teacherHandler.CreateTeacher)
			teachers.PUT("/:id", teacherHandler.UpdateTeacher)
			teachers.DELETE("/:id", teacherHandler.DeleteTeacher)
		}

		students := v1.Group("/students")
		{
			students.GET("", studentHandler.GetAllStudents)
			students.GET("/:id", studentHandler.GetStudent)
			students.POST("", studentHandler.CreateStudent)
			students.PUT("/:id", studentHandler.UpdateStudent)
			students.DELETE("/:id", studentHandler.DeleteStudent)
		}

		subjects := v1.Group("/subjects")
		{
			subjects.GET("", subjectHandler.GetAllSubjects)
			subjects.GET("/:id", subjectHandler.GetSubject)
			subjects.POST("", subjectHandler.CreateSubject)
			subjects.PUT("/:id", subjectHandler.UpdateSubject)
			subjects.DELETE("/:id", subjectHandler.DeleteSubject)
		}
	}

	return router
}

