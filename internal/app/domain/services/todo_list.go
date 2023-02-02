package services

import (
	"errors"
	"hexa-example-go/internal/app/domain/adapters"
	"hexa-example-go/internal/app/domain/entities"

	"go.uber.org/zap"
)

type todoListService struct {
	logger zap.Logger
	repo   adapters.TodoListRepository
}

func NewTodoListService(logger zap.Logger, repo adapters.TodoListRepository) adapters.TodoListService {
	return &todoListService{
		logger: logger,
		repo:   repo,
	}
}

func (s *todoListService) CreateTodoList(dto adapters.CreateTodoListDTO) (*entities.TodoList, error) {
	newName := dto.Name

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
