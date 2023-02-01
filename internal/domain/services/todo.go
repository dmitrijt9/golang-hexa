package services

import (
	"errors"
	"hexa-example-go/internal/domain/adapters"
	"hexa-example-go/internal/domain/entities"

	"go.uber.org/zap"
)

type TodoService struct {
	logger   zap.Logger
	repo     adapters.TodoRepository
	repoList adapters.TodoListRepository
}

func NewTodoService(logger zap.Logger, repo adapters.TodoRepository, repoList adapters.TodoListRepository) *TodoService {
	return &TodoService{
		logger:   logger,
		repo:     repo,
		repoList: repoList,
	}
}

type createTodoDTO struct {
	title       string
	description string
	status      entities.TodoStatus
	listName    string
}

func (s *TodoService) CreateTodo(dto createTodoDTO) (*entities.Todo, error) {
	newTitle := dto.title
	newDescription := dto.description
	status := dto.status
	listName := dto.listName

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

	todoToSave := adapters.TodoToSave{
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
