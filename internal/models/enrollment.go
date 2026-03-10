package models

import "time"

type Enrollment struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	StudentID      uint      `json:"studentId" gorm:"column:studentId;not null"`
	SubjectID      uint      `json:"subjectId" gorm:"column:subjectId;not null"`
	EnrollmentDate time.Time `json:"enrollmentDate" gorm:"column:enrollmentDate"`
	Grade          string    `json:"grade" gorm:"column:grade"`
	Student        Student   `json:"student" gorm:"foreignKey:StudentID"`
	Subject        Subject   `json:"subject" gorm:"foreignKey:SubjectID"`
}

func (Enrollment) TableName() string {
	return "enrollments"
}

type CreateEnrollmentRequest struct {
	StudentID      uint      `json:"studentId" binding:"required"`
	SubjectID      uint      `json:"subjectId" binding:"required"`
	EnrollmentDate time.Time `json:"enrollmentDate" binding:"required"`
	Grade          string    `json:"grade"`
}

type UpdateEnrollmentRequest struct {
	Grade string `json:"grade" binding:"required"`
}

func (r *CreateEnrollmentRequest) ToEnrollment() *Enrollment {
	return &Enrollment{
		StudentID:      r.StudentID,
		SubjectID:      r.SubjectID,
		EnrollmentDate: r.EnrollmentDate,
		Grade:          r.Grade,
	}
}

