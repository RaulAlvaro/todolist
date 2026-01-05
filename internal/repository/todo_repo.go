package repository

import (
	"context"
	"todolist/internal/domain"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

// NewTodoRepository devuelve la INTERFAZ, no el puntero a la estructura privada
func NewTodoRepository(db *gorm.DB) domain.TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) GetAll(ctx context.Context) ([]domain.Todo, error) {
	var todos []domain.Todo
	// Usamos WithContext(ctx) para que la consulta respete cancelaciones o timeouts
	err := r.db.WithContext(ctx).Find(&todos).Error
	return todos, err
}

func (r *todoRepository) Create(ctx context.Context, todo *domain.Todo) error {
	return r.db.WithContext(ctx).Create(todo).Error
}

func (r *todoRepository) GetByID(ctx context.Context, id uint) (*domain.Todo, error) {
	todo := &domain.Todo{}
	err := r.db.WithContext(ctx).First(todo, id).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *todoRepository) Update(ctx context.Context, todo *domain.Todo) error {
	return r.db.WithContext(ctx).Save(todo).Error
}

func (r *todoRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&domain.Todo{}, id).Error
}
