//go:build wireinject
// +build wireinject

package main

import (
	"go-boilerplate/src/http_app"
	"go-boilerplate/src/service/class_service"
	"go-boilerplate/src/service/subject_service"
	"go-boilerplate/src/storage"

	"github.com/google/wire"
)

func InitilizeAppRequirements(config storage.DbConfig) *http_app.App {
	wire.Build(
		storage.NewDB,
		storage.NewClassesRepo,
		class_service.NewClassService,
		storage.NewSubjectsRepo,
		subject_service.NewSubjectService,
		http_app.Initialize,
		wire.Bind(new(class_service.ClassesRepo), new(*storage.ClassesRepo)),
		wire.Bind(new(http_app.ClassService), new(*class_service.ClassService)),
		wire.Bind(new(subject_service.SubjectsRepo), new(*storage.SubjectsRepo)),
		wire.Bind(new(http_app.SubjectService), new(*subject_service.SubjectService)),
	)

	return &http_app.App{}
}
