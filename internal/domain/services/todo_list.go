package services

import (
	"errors"
	"hexa-example-go/internal/domain/adapters"
	"hexa-example-go/internal/domain/entities"

	"go.uber.org/zap"
)

type TodoListService struct {
	logger zap.Logger
	repo   adapters.TodoListRepository
}

func NewTodoListService(logger zap.Logger, repo adapters.TodoListRepository) *TodoListService {
	return &TodoListService{
		logger: logger,
		repo:   repo,
	}
}

type CreateTodoListDTO struct {
	name string
}

func (s *TodoListService) CreateTodoList(dto CreateTodoListDTO) (*entities.TodoList, error) {
	newName := dto.name

	if newName == "" {
		return nil, errors.New("name of todo list cannot be empty")
	}

	found, err := s.repo.GetByName(newName)
	if err != nil {
		return nil, err
	}

	if found != nil {
		return nil, errors.New("todo list with such name already exist")
	}

	todolistToSave := adapters.TodoListToSave{
		Name: newName,
	}
	newTodoList, err := s.repo.Save(todolistToSave)
	if err != nil {
		return nil, err
	}

	return newTodoList, err
}
