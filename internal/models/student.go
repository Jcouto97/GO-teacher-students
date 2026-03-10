package models

import "time"

type Student struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name" gorm:"not null" binding:"required"`
	Email          string    `json:"email" gorm:"unique;not null" binding:"required,email"`
	EnrollmentDate time.Time `json:"enrollmentDate" gorm:"column:enrollmentDate"`
}

func (Student) TableName() string {
	return "students"
}

type CreateStudentRequest struct {
	Name           string    `json:"name" binding:"required"`
	Email          string    `json:"email" binding:"required,email"`
	EnrollmentDate time.Time `json:"enrollmentDate" binding:"required"`
}

type UpdateStudentRequest struct {
	Name           string    `json:"name" binding:"required"`
	Email          string    `json:"email" binding:"required,email"`
	EnrollmentDate time.Time `json:"enrollmentDate" binding:"required"`
}

func (r *CreateStudentRequest) ToStudent() *Student {
	return &Student{
		Name:           r.Name,
		Email:          r.Email,
		EnrollmentDate: r.EnrollmentDate,
	}
}

