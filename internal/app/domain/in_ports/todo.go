package in_ports

import "hexa-example-go/internal/app/domain/entities"

type CreateTodoDTO struct {
	Title       string
	Description string
	Status      entities.TodoStatus
	ListName    string
}

type TodoService interface {
	CreateTodo(dto CreateTodoDTO) (*entities.Todo, error)
}
