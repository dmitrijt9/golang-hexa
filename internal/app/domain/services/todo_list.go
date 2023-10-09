package services

import (
	"errors"
	"fmt"
	"hexa-example-go/internal/app/domain/entities"
	"hexa-example-go/internal/app/domain/in_ports"
	"hexa-example-go/internal/app/domain/out_ports"
	"hexa-example-go/internal/logger"
)

type todoListService struct {
	logger      logger.Logger
	repo        out_ports.TodoListRepository
	mailService out_ports.MailService
}

func NewTodoListService(logger logger.Logger, repo out_ports.TodoListRepository) in_ports.TodoListService {
	return &todoListService{
		logger: logger,
		repo:   repo,
		// TODO: pass mail service dependency
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

	mailMessage := fmt.Sprintf("Your new todo list '%s' was successfully created! :)", newName)
	s.mailService.SendMessage("me@me.com", "system@hexa-example-go.com", "New Todo List created!", mailMessage)

	return newTodoList, err
}
