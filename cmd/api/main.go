package main

import (
	"fmt"
	"goProject/internal/config"
	"goProject/internal/handler"
	"goProject/internal/repository"
	"goProject/internal/service"
	"goProject/pkg/database"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.NewMySQLConnection(cfg.Database.GetDSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// dependency Injection
	teacherRepo := repository.NewTeacherRepository(db)
	teacherService := service.NewTeacherService(teacherRepo)
	teacherHandler := handler.NewTeacherHandler(teacherService)

	studentRepo := repository.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepo)
	studentHandler := handler.NewStudentHandler(studentService)

	subjectRepo := repository.NewSubjectRepository(db)
	subjectService := service.NewSubjectService(subjectRepo)
	subjectHandler := handler.NewSubjectHandler(subjectService)

	enrollmentRepo := repository.NewEnrollmentRepository(db)
	enrollmentService := service.NewEnrollmentService(enrollmentRepo, studentRepo, subjectRepo)
	enrollmentHandler := handler.NewEnrollmentHandler(enrollmentService)

	router := handler.SetupRouter(teacherHandler, studentHandler, subjectHandler, enrollmentHandler)

	// start server
	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Server starting on http://localhost%s", serverAddr)
	
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

