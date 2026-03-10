package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(teacherHandler *TeacherHandler, studentHandler *StudentHandler, subjectHandler *SubjectHandler, enrollmentHandler *EnrollmentHandler) *gin.Engine {
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
			students.GET("/:id/subjects", enrollmentHandler.GetStudentEnrollments)
			students.POST("", studentHandler.CreateStudent)
			students.PUT("/:id", studentHandler.UpdateStudent)
			students.DELETE("/:id", studentHandler.DeleteStudent)
		}

		subjects := v1.Group("/subjects")
		{
			subjects.GET("", subjectHandler.GetAllSubjects)
			subjects.GET("/:id", subjectHandler.GetSubject)
			subjects.GET("/:id/students", enrollmentHandler.GetSubjectEnrollments)
			subjects.POST("", subjectHandler.CreateSubject)
			subjects.PUT("/:id", subjectHandler.UpdateSubject)
			subjects.DELETE("/:id", subjectHandler.DeleteSubject)
		}

		enrollments := v1.Group("/enrollments")
		{
			enrollments.GET("", enrollmentHandler.GetAllEnrollments)
			enrollments.GET("/:id", enrollmentHandler.GetEnrollment)
			enrollments.POST("", enrollmentHandler.CreateEnrollment)
			enrollments.PUT("/:id", enrollmentHandler.UpdateEnrollment)
			enrollments.DELETE("/:id", enrollmentHandler.DeleteEnrollment)
		}
	}

	return router
}

