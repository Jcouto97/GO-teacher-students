package models

type Subject struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null" binding:"required"`
	Description string `json:"description"`
	TeacherID   uint   `json:"teacherId" gorm:"column:teacherId;not null"`
	Teacher     Teacher `json:"teacher" gorm:"foreignKey:TeacherID"`
}

func (Subject) TableName() string {
	return "subjects"
}

type CreateSubjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	TeacherID   uint   `json:"teacherId" binding:"required"`
}

type UpdateSubjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	TeacherID   uint   `json:"teacherId" binding:"required"`
}

func (r *CreateSubjectRequest) ToSubject() *Subject {
	return &Subject{
		Name:        r.Name,
		Description: r.Description,
		TeacherID:   r.TeacherID,
	}
}

