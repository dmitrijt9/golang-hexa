package services

import (
	"errors"
	"hexa-example-go/internal/app/domain/entities"
	"hexa-example-go/internal/app/domain/in_ports"
	"hexa-example-go/internal/app/domain/out_ports"

	"go.uber.org/zap"
)

type TodoService struct {
	logger   zap.Logger
	repo     out_ports.TodoRepository
	repoList out_ports.TodoListRepository
}

func NewTodoService(logger zap.Logger, repo out_ports.TodoRepository, repoList out_ports.TodoListRepository) *TodoService {
	return &TodoService{
		logger:   logger,
		repo:     repo,
		repoList: repoList,
	}
}

func (s *TodoService) CreateTodo(dto in_ports.CreateTodoDTO) (*entities.Todo, error) {
	newTitle := dto.Title
	newDescription := dto.Description
	status := dto.Status
	listName := dto.ListName

	if newTitle == "" {
		return nil, errors.New("title of the todo cannot be empty")
	}

	if status == "" {
		status = entities.Open
	}

	if listName == "" {
		listName = "Inbox"
	}

	todoList, err := s.repoList.GetByName(listName)
	if err != nil {
		return nil, err
	}

	todoToSave := out_ports.TodoToSave{
		Title:       newTitle,
		Description: newDescription,
		Status:      status,
		List:        todoList,
	}
	newTodo, err := s.repo.Save(todoToSave)
	if err != nil {
		return nil, err
	}

	return newTodo, err
}
