package out_ports

import "hexa-example-go/internal/app/domain/entities"

type TodoListToSave struct {
	Name string
}

type TodoListRepository interface {
	Save(toSave TodoListToSave) (*entities.TodoList, error)
	GetByName(name string) (*entities.TodoList, error)
	List() ([]*entities.TodoList, error)
}
