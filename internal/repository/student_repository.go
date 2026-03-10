package repository

import (
	"goProject/internal/models"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FindAll() ([]models.Student, error)
	FindByID(id uint) (*models.Student, error)
	Create(student *models.Student) error
	Update(student *models.Student) error
	Delete(id uint) error
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db: db}
}

func (r *studentRepository) FindAll() ([]models.Student, error) {
	var students []models.Student
	err := r.db.Find(&students).Error
	return students, err
}

func (r *studentRepository) FindByID(id uint) (*models.Student, error) {
	var student models.Student
	err := r.db.First(&student, id).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepository) Create(student *models.Student) error {
	return r.db.Create(student).Error
}

func (r *studentRepository) Update(student *models.Student) error {
	return r.db.Save(student).Error
}

func (r *studentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Student{}, id).Error
}

