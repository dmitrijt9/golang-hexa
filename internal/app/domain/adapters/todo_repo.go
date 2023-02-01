package adapters

import "hexa-example-go/internal/app/domain/entities"

type TodoToSave struct {
	Title       string
	Description string
	Status      entities.TodoStatus
	List        *entities.TodoList
}

type TodoRepository interface {
	Save(toSave TodoToSave) (*entities.Todo, error)
	List() ([]*entities.Todo, error)
}
