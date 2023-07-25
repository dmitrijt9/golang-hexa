package graph

import (
	"hexa-example-go/internal/app/domain/in_ports"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoListService in_ports.TodoListService
}
