package entities

type TodoStatus string

const (
	Open       TodoStatus = "Open"
	InProgress TodoStatus = "InProgress"
	Done       TodoStatus = "Done"
)

type Todo struct {
	Id          int
	Title       string
	Description string
	Status      TodoStatus
	List        *TodoList
}
