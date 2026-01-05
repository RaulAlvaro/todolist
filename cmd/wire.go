//go:build wireinject
// +build wireinject

package main

import (
	"todolist/config"
	"todolist/internal/db"
	"todolist/internal/handler"
	"todolist/internal/repository"
	"todolist/internal/service"

	"github.com/google/wire"
)

func initializeApp() (*App, error) {
	wire.Build(
		config.Load,
		db.ProvideDatabase,
		repository.NewTodoRepository,
		service.NewTodoService,
		handler.NewTodoHandler,
		newApp, // Esta función la definiremos en el main.go
	)
	return nil, nil
}
