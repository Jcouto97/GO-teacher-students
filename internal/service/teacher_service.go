package service

import (
	"errors"
	"goProject/internal/models"
	"goProject/internal/repository"

	"gorm.io/gorm"
)

type TeacherService interface {
	GetAllTeachers() ([]models.Teacher, error)
	GetTeacherByID(id uint) (*models.Teacher, error)
	CreateTeacher(req *models.CreateTeacherRequest) (*models.Teacher, error)
	UpdateTeacher(id uint, req *models.UpdateTeacherRequest) (*models.Teacher, error)
	DeleteTeacher(id uint) error
}

type teacherService struct {
	repo repository.TeacherRepository
}

func NewTeacherService(repo repository.TeacherRepository) TeacherService {
	return &teacherService{repo: repo}
}

func (s *teacherService) GetAllTeachers() ([]models.Teacher, error) {
	return s.repo.FindAll()
}

func (s *teacherService) GetTeacherByID(id uint) (*models.Teacher, error) {
	teacher, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("teacher not found")
		}
		return nil, err
	}
	return teacher, nil
}

func (s *teacherService) CreateTeacher(req *models.CreateTeacherRequest) (*models.Teacher, error) {
	teacher := req.ToTeacher()
	
	if err := s.repo.Create(teacher); err != nil {
		return nil, err
	}
	
	return teacher, nil
}

func (s *teacherService) UpdateTeacher(id uint, req *models.UpdateTeacherRequest) (*models.Teacher, error) {
	// Check if teacher exists
	teacher, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("teacher not found")
		}
		return nil, err
	}

	teacher.Name = req.Name
	teacher.Email = req.Email
	teacher.SubjectSpecialization = req.SubjectSpecialization

	if err := s.repo.Update(teacher); err != nil {
		return nil, err
	}

	return teacher, nil
}

func (s *teacherService) DeleteTeacher(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("teacher not found")
		}
		return err
	}

	return s.repo.Delete(id)
}

