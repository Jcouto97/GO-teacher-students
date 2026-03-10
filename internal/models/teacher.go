package models

type Teacher struct {
	ID                    uint   `json:"id" gorm:"primaryKey"`
	Name                  string `json:"name" gorm:"not null" binding:"required"`
	Email                 string `json:"email" gorm:"unique;not null" binding:"required,email"`
	SubjectSpecialization string `json:"subjectSpecialization" gorm:"column:subjectSpecialization"`
}

// specifies the table name for GORM
func (Teacher) TableName() string {
	return "teachers"
}

type CreateTeacherRequest struct {
	Name                  string `json:"name" binding:"required"`
	Email                 string `json:"email" binding:"required,email"`
	SubjectSpecialization string `json:"subjectSpecialization"`
}

type UpdateTeacherRequest struct {
	Name                  string `json:"name" binding:"required"`
	Email                 string `json:"email" binding:"required,email"`
	SubjectSpecialization string `json:"subjectSpecialization"`
}

func (r *CreateTeacherRequest) ToTeacher() *Teacher {
	return &Teacher{
		Name:                  r.Name,
		Email:                 r.Email,
		SubjectSpecialization: r.SubjectSpecialization,
	}
}

