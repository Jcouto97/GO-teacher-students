package service

import (
	"errors"
	"goProject/internal/models"
	"goProject/internal/repository"

	"gorm.io/gorm"
)

type EnrollmentService interface {
	GetAllEnrollments() ([]models.Enrollment, error)
	GetEnrollmentByID(id uint) (*models.Enrollment, error)
	GetStudentEnrollments(studentID uint) ([]models.Enrollment, error)
	GetSubjectEnrollments(subjectID uint) ([]models.Enrollment, error)
	CreateEnrollment(req *models.CreateEnrollmentRequest) (*models.Enrollment, error)
	UpdateEnrollment(id uint, req *models.UpdateEnrollmentRequest) (*models.Enrollment, error)
	DeleteEnrollment(id uint) error
}

type enrollmentService struct {
	repo        repository.EnrollmentRepository
	studentRepo repository.StudentRepository
	subjectRepo repository.SubjectRepository
}

func NewEnrollmentService(
	repo repository.EnrollmentRepository,
	studentRepo repository.StudentRepository,
	subjectRepo repository.SubjectRepository,
) EnrollmentService {
	return &enrollmentService{
		repo:        repo,
		studentRepo: studentRepo,
		subjectRepo: subjectRepo,
	}
}

func (s *enrollmentService) GetAllEnrollments() ([]models.Enrollment, error) {
	return s.repo.FindAll()
}

func (s *enrollmentService) GetEnrollmentByID(id uint) (*models.Enrollment, error) {
	enrollment, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("enrollment not found")
		}
		return nil, err
	}
	return enrollment, nil
}

func (s *enrollmentService) GetStudentEnrollments(studentID uint) ([]models.Enrollment, error) {
	_, err := s.studentRepo.FindByID(studentID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("student not found")
		}
		return nil, err
	}
	return s.repo.FindByStudentID(studentID)
}

func (s *enrollmentService) GetSubjectEnrollments(subjectID uint) ([]models.Enrollment, error) {
	_, err := s.subjectRepo.FindByID(subjectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("subject not found")
		}
		return nil, err
	}
	return s.repo.FindBySubjectID(subjectID)
}

func (s *enrollmentService) CreateEnrollment(req *models.CreateEnrollmentRequest) (*models.Enrollment, error) {
	_, err := s.studentRepo.FindByID(req.StudentID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("student not found")
		}
		return nil, err
	}

	_, err = s.subjectRepo.FindByID(req.SubjectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("subject not found")
		}
		return nil, err
	}

	exists, err := s.repo.CheckExists(req.StudentID, req.SubjectID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("student already enrolled in this subject")
	}

	enrollment := req.ToEnrollment()
	if err := s.repo.Create(enrollment); err != nil {
		return nil, err
	}

	return s.repo.FindByID(enrollment.ID)
}

func (s *enrollmentService) UpdateEnrollment(id uint, req *models.UpdateEnrollmentRequest) (*models.Enrollment, error) {
	enrollment, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("enrollment not found")
		}
		return nil, err
	}

	enrollment.Grade = req.Grade

	if err := s.repo.Update(enrollment); err != nil {
		return nil, err
	}

	return s.repo.FindByID(enrollment.ID)
}

func (s *enrollmentService) DeleteEnrollment(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("enrollment not found")
		}
		return err
	}

	return s.repo.Delete(id)
}

