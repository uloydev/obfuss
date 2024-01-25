package services

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
)

type TodoService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewTodoService(db *gorm.DB, logger *zap.Logger) *TodoService {
	return &TodoService{
		db,
		logger.With(zap.String("type", "TodoService")),
	}
}

func (s *TodoService) GetTodos() ([]entities.Todo, error) {
	var todos []entities.Todo

	if err := s.db.Find(&todos).Error; err != nil {
		return nil, errors.New("error when get todos")
	}

	return todos, nil
}

func (s *TodoService) GetTodoByID(id uint) (*entities.Todo, error) {
	var todo entities.Todo

	if err := s.db.First(&todo, id).Error; err != nil {
		return nil, errors.New("error when get todo")
	}

	return &todo, nil
}

func (s *TodoService) Create(todo *entities.Todo) error {
	if err := s.db.Create(todo).Error; err != nil {
		return errors.New("error when create todo")
	}

	return nil
}

func (s *TodoService) Update(todo *entities.Todo) error {
	if err := s.db.Save(todo).Error; err != nil {
		return errors.New("error when update todo")
	}

	return nil
}

func (s *TodoService) Delete(id uint) error {
	if err := s.db.Delete(id).Error; err != nil {
		return errors.New("error when delete todo")
	}

	return nil
}
