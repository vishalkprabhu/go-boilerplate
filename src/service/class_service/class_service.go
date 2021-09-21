package class_service

import "go-boilerplate/src/model"

type ClassService struct {
	classRepo ClassesRepo
}

type ClassesRepo interface {
	Create(class model.Class) (*model.Class, error)
	GetAll() ([]model.Class, error)
}

func NewClassService(classRepo ClassesRepo) *ClassService {
	return &ClassService{classRepo: classRepo}
}

func (s *ClassService) List() []string {

	return []string{"1", "2"}

}

func (s *ClassService) Create() bool {
	return true
}
