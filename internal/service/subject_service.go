package service

import (
	"errors"
	"goProject/internal/models"
	"goProject/internal/repository"

	"gorm.io/gorm"
)

type SubjectService interface {
	GetAllSubjects() ([]models.Subject, error)
	GetSubjectByID(id uint) (*models.Subject, error)
	CreateSubject(req *models.CreateSubjectRequest) (*models.Subject, error)
	UpdateSubject(id uint, req *models.UpdateSubjectRequest) (*models.Subject, error)
	DeleteSubject(id uint) error
}

type subjectService struct {
	repo repository.SubjectRepository
}

func NewSubjectService(repo repository.SubjectRepository) SubjectService {
	return &subjectService{repo: repo}
}

func (s *subjectService) GetAllSubjects() ([]models.Subject, error) {
	return s.repo.FindAll()
}

func (s *subjectService) GetSubjectByID(id uint) (*models.Subject, error) {
	subject, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("subject not found")
		}
		return nil, err
	}
	return subject, nil
}

func (s *subjectService) CreateSubject(req *models.CreateSubjectRequest) (*models.Subject, error) {
	subject := req.ToSubject()
	
	if err := s.repo.Create(subject); err != nil {
		return nil, err
	}
	
	return subject, nil
}

func (s *subjectService) UpdateSubject(id uint, req *models.UpdateSubjectRequest) (*models.Subject, error) {
	subject, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("subject not found")
		}
		return nil, err
	}

	subject.Name = req.Name
	subject.Description = req.Description
	subject.TeacherID = req.TeacherID

	if err := s.repo.Update(subject); err != nil {
		return nil, err
	}

	return subject, nil
}

func (s *subjectService) DeleteSubject(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("subject not found")
		}
		return err
	}

	return s.repo.Delete(id)
}

