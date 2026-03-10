package service

import (
	"errors"
	"goProject/internal/models"
	"goProject/internal/repository"

	"gorm.io/gorm"
)

type StudentService interface {
	GetAllStudents() ([]models.Student, error)
	GetStudentByID(id uint) (*models.Student, error)
	CreateStudent(req *models.CreateStudentRequest) (*models.Student, error)
	UpdateStudent(id uint, req *models.UpdateStudentRequest) (*models.Student, error)
	DeleteStudent(id uint) error
}

type studentService struct {
	repo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{repo: repo}
}

func (s *studentService) GetAllStudents() ([]models.Student, error) {
	return s.repo.FindAll()
}

func (s *studentService) GetStudentByID(id uint) (*models.Student, error) {
	student, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("student not found")
		}
		return nil, err
	}
	return student, nil
}

func (s *studentService) CreateStudent(req *models.CreateStudentRequest) (*models.Student, error) {
	student := req.ToStudent()
	
	if err := s.repo.Create(student); err != nil {
		return nil, err
	}
	
	return student, nil
}

func (s *studentService) UpdateStudent(id uint, req *models.UpdateStudentRequest) (*models.Student, error) {
	student, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("student not found")
		}
		return nil, err
	}

	student.Name = req.Name
	student.Email = req.Email
	student.EnrollmentDate = req.EnrollmentDate

	if err := s.repo.Update(student); err != nil {
		return nil, err
	}

	return student, nil
}

func (s *studentService) DeleteStudent(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("student not found")
		}
		return err
	}

	return s.repo.Delete(id)
}

