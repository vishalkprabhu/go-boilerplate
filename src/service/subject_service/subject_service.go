package subject_service

import "go-boilerplate/src/model"

type SubjectService struct {
	subjectRepo SubjectsRepo
}

type SubjectsRepo interface {
	Create(subject model.Subject) (*model.Subject, error)
	GetAll() ([]model.Subject, error)
}

func NewSubjectService(subjectRepo SubjectsRepo) *SubjectService {
	return &SubjectService{subjectRepo: subjectRepo}
}

func (s *SubjectService) List() []string {

	return []string{"1", "2"}

}

func (s *SubjectService) Create() bool {
	return true
}
