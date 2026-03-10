package repository

import (
	"goProject/internal/models"

	"gorm.io/gorm"
)

type TeacherRepository interface {
	FindAll() ([]models.Teacher, error)
	FindByID(id uint) (*models.Teacher, error)
	Create(teacher *models.Teacher) error
	Update(teacher *models.Teacher) error
	Delete(id uint) error
}

type teacherRepository struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) TeacherRepository {
	return &teacherRepository{db: db}
}

func (r *teacherRepository) FindAll() ([]models.Teacher, error) {
	var teachers []models.Teacher
	err := r.db.Find(&teachers).Error
	return teachers, err
}

func (r *teacherRepository) FindByID(id uint) (*models.Teacher, error) {
	var teacher models.Teacher
	err := r.db.First(&teacher, id).Error
	if err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (r *teacherRepository) Create(teacher *models.Teacher) error {
	return r.db.Create(teacher).Error
}

func (r *teacherRepository) Update(teacher *models.Teacher) error {
	return r.db.Save(teacher).Error
}

func (r *teacherRepository) Delete(id uint) error {
	return r.db.Delete(&models.Teacher{}, id).Error
}

