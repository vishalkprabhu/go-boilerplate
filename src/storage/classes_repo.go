package storage

import (
	"go-boilerplate/src/model"

	"github.com/jinzhu/gorm"
)

type ClassesRepo struct {
	db *gorm.DB
}

func NewClassesRepo(db *gorm.DB) *ClassesRepo {
	return &ClassesRepo{
		db: db,
	}
}

func (r *ClassesRepo) Create(class model.Class) (*model.Class, error) {

	return nil, nil
}

func (r *ClassesRepo) GetAll() ([]model.Class, error) {

	return nil, nil
}
