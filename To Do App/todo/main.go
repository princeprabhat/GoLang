package todo

import (
	"time"
)

// type todoType interface {
// 	// AddTodo adds a new todo item to the list
// 	AddTodo(todo string) error
// 	// RemoveTodo removes a todo item from the list
// 	RemoveTodo(todo string) error
// 	// ListTodo lists all todo items
// 	ListTodo() []string
// }

type todo struct {
	ID        int
	Task      string
	Completed bool
	Added     time.Time
}

type todoList struct {
	todoStore []todo
}

func NewTodoService() *todoList {
	return &todoList{}
}
