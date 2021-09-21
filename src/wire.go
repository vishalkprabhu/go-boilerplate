//+build wireinject

package main

import (
	"go-boilerplate/src/http_app"
	"go-boilerplate/src/service/class_service"
	"go-boilerplate/src/storage"

	"github.com/google/wire"
)

func InitilizeAppRequirements(config storage.DbConfig) *http_app.App {
	wire.Build(
		storage.NewDB,
		storage.NewClassesRepo,
		class_service.NewClassService,
		http_app.Initialize,
		wire.Bind(new(class_service.ClassesRepo), new(*storage.ClassesRepo)),
		wire.Bind(new(http_app.ClassService), new(*class_service.ClassService)),
	)

	return &http_app.App{}
}
