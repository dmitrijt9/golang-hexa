package adapters

import "hexa-example-go/internal/app/domain/entities"

// This could be also in a separate directory (eg. /services/dto/)
// I like it here, because it is right above the actual usage.
type CreateTodoListDTO struct {
	Name string
}

type TodoListService interface {
	CreateTodoList(dto CreateTodoListDTO) (*entities.TodoList, error)
}
