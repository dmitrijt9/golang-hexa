package services

import (
	"errors"
	"hexa-example-go/internal/app/domain/adapters"
	"hexa-example-go/internal/app/domain/entities"

	"go.uber.org/zap"
)

type TodoListService interface {
	CreateTodoList(dto createTodoListDTO) (*entities.TodoList, error)
}

type todoListService struct {
	logger zap.Logger
	repo   adapters.TodoListRepository
}

func NewTodoListService(logger zap.Logger, repo adapters.TodoListRepository) TodoListService {
	return &todoListService{
		logger: logger,
		repo:   repo,
	}
}

// This could be also in a separate directory (eg. /services/dto/)
// I like it here, because it is right above the actual usage.
type createTodoListDTO struct {
	name string
}

func (s *todoListService) CreateTodoList(dto createTodoListDTO) (*entities.TodoList, error) {
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
