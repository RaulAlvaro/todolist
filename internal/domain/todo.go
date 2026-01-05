package domain

import (
	"context"

	"gorm.io/gorm"
)

// Todo es nuestra entidad de base de datos
type Todo struct {
	gorm.Model
	Content string `json:"content"`
	Status  bool   `json:"status"`
}

// TodoRepository define qué puede hacer la base de datos con los Todos.
// Usamos interfaces para que el Service no sepa si usamos Postgres o MySQL.
type TodoRepository interface {
	Create(ctx context.Context, todo *Todo) error
	GetByID(ctx context.Context, id uint) (*Todo, error)
	GetAll(ctx context.Context) ([]Todo, error)
	Update(ctx context.Context, todo *Todo) error
	Delete(ctx context.Context, id uint) error
}
