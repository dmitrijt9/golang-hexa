package in_ports

import (
	"hexa-example-go/internal/app/domain/entities"
)

type CreateTodoListDTO struct {
	Name string
}

type TodoListService interface {
	CreateTodoList(dto CreateTodoListDTO) (*entities.TodoList, error)
}
