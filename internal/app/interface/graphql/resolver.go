package graph

import (
	"hexa-example-go/internal/app/domain/out_ports"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoListService out_ports.TodoListService
}
