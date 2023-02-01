package graph

import (
	"hexa-example-go/internal/app/domain/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoListService services.TodoListService
}
