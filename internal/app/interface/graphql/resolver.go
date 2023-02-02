package graph

import (
	"hexa-example-go/internal/app/domain/adapters"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoListService adapters.TodoListService
}
