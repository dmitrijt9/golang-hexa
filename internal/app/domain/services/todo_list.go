package services

import (
	"errors"
	"hexa-example-go/internal/app/domain/entities"
	"hexa-example-go/internal/app/domain/in_ports"
	"hexa-example-go/internal/app/domain/out_ports"

	"go.uber.org/zap"
)

type todoListService struct {
	logger zap.Logger
	repo   out_ports.TodoListRepository
}

func NewTodoListService(logger zap.Logger, repo out_ports.TodoListRepository) in_ports.TodoListService {
	return &todoListService{
		logger: logger,
		repo:   repo,
	}
}

func (s *todoListService) CreateTodoList(dto in_ports.CreateTodoListDTO) (*entities.TodoList, error) {
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

	todolistToSave := out_ports.TodoListToSave{
		Name: newName,
	}
	newTodoList, err := s.repo.Save(todolistToSave)
	if err != nil {
		return nil, err
	}

	return newTodoList, err
}
