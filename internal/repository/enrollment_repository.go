package repository

import (
	"goProject/internal/models"

	"gorm.io/gorm"
)

type EnrollmentRepository interface {
	FindAll() ([]models.Enrollment, error)
	FindByID(id uint) (*models.Enrollment, error)
	FindByStudentID(studentID uint) ([]models.Enrollment, error)
	FindBySubjectID(subjectID uint) ([]models.Enrollment, error)
	Create(enrollment *models.Enrollment) error
	Update(enrollment *models.Enrollment) error
	Delete(id uint) error
	CheckExists(studentID, subjectID uint) (bool, error)
}

type enrollmentRepository struct {
	db *gorm.DB
}

func NewEnrollmentRepository(db *gorm.DB) EnrollmentRepository {
	return &enrollmentRepository{db: db}
}

func (r *enrollmentRepository) FindAll() ([]models.Enrollment, error) {
	var enrollments []models.Enrollment
	err := r.db.Preload("Student").Preload("Subject.Teacher").Find(&enrollments).Error
	return enrollments, err
}

func (r *enrollmentRepository) FindByID(id uint) (*models.Enrollment, error) {
	var enrollment models.Enrollment
	err := r.db.Preload("Student").Preload("Subject.Teacher").First(&enrollment, id).Error
	if err != nil {
		return nil, err
	}
	return &enrollment, nil
}

func (r *enrollmentRepository) FindByStudentID(studentID uint) ([]models.Enrollment, error) {
	var enrollments []models.Enrollment
	err := r.db.Preload("Subject.Teacher").Where("studentId = ?", studentID).Find(&enrollments).Error
	return enrollments, err
}

func (r *enrollmentRepository) FindBySubjectID(subjectID uint) ([]models.Enrollment, error) {
	var enrollments []models.Enrollment
	err := r.db.Preload("Student").Where("subjectId = ?", subjectID).Find(&enrollments).Error
	return enrollments, err
}

func (r *enrollmentRepository) Create(enrollment *models.Enrollment) error {
	return r.db.Create(enrollment).Error
}

func (r *enrollmentRepository) Update(enrollment *models.Enrollment) error {
	return r.db.Save(enrollment).Error
}

func (r *enrollmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Enrollment{}, id).Error
}

func (r *enrollmentRepository) CheckExists(studentID, subjectID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.Enrollment{}).
		Where("studentId = ? AND subjectId = ?", studentID, subjectID).
		Count(&count).Error
	return count > 0, err
}

