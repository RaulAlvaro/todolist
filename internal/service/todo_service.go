package service

import (
	"context"
	"fmt"
	"todolist/internal/domain"
	"todolist/internal/repository"

	"gorm.io/gorm"
)

type TodoService struct {
	db       *gorm.DB
	todoRepo domain.TodoRepository
}

func NewTodoService(db *gorm.DB, repo domain.TodoRepository) *TodoService {
	return &TodoService{
		db:       db,
		todoRepo: repo,
	}
}

func (s *TodoService) GetAll(ctx context.Context) ([]domain.Todo, error) {
	return s.todoRepo.GetAll(ctx)
}

func (s *TodoService) Create(ctx context.Context, todo *domain.Todo) error {
	if todo.Content == "" {
		return fmt.Errorf("el contenido del todo no puede estar vacío")
	}
	return s.todoRepo.Create(ctx, todo)
}

func (s *TodoService) CreateWithAudit(ctx context.Context, todo *domain.Todo) error {
	// Iniciamos la transacción usando el puntero s.db que inyectamos con Wire
	return s.db.Transaction(func(tx *gorm.DB) error {

		// REGLA DE ORO: Dentro de la transacción,
		// debemos crear una instancia del repositorio que use 'tx' (la transacción)
		// y no 's.db' (la conexión global).
		txRepo := repository.NewTodoRepository(tx)

		// 1. Intentamos crear el Todo
		if err := txRepo.Create(ctx, todo); err != nil {
			return err // Al retornar error, GORM hace ROLLBACK automático
		}

		// 2. Simulamos otra operación (ej. Auditoría)
		// Si aquí retornaras un error, el Todo de arriba NO se guardaría
		// err := auditRepo.Log(ctx, "Se creó un nuevo todo")

		return nil // Si retorna nil, GORM hace COMMIT automático
	})
}

func (s *TodoService) GetById(ctx context.Context, id uint) (*domain.Todo, error) {
	if id == 0 {
		return nil, fmt.Errorf("ID no proporcionado")
	}

	todo, err := s.todoRepo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

// func (s *TodoService) Update(ctx context.Context, todo *domain.Todo) error {
// 	if
// }
