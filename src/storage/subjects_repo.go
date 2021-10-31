package storage

import (
	"go-boilerplate/src/model"

	"github.com/jinzhu/gorm"
)

type SubjectsRepo struct {
	db *gorm.DB
}

func NewSubjectsRepo(db *gorm.DB) *SubjectsRepo {
	return &SubjectsRepo{
		db: db,
	}
}

func (r *SubjectsRepo) Create(subject model.Subject) (*model.Subject, error) {

	return nil, nil
}

func (r *SubjectsRepo) GetAll() ([]model.Subject, error) {

	return nil, nil
}
