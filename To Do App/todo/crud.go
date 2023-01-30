package todo

import (
	"fmt"
	"strings"
	"time"
)

// AddTodo adds a new todo item to the list
func (t *todoList) AddTodo(newTodo []string) {
	t.loadFromJson()

	newTodoIns := todo{
		ID:        t.newTodoId(),
		Task:      strings.Join(newTodo, " "),
		Completed: false,
		Added:     time.Now(),
	}
	t.todoStore = append(t.todoStore, newTodoIns)

	// Save the todo list to a json file
	t.saveToJson()
	fmt.Println("Added todo:", newTodo)
}

// RemoveTodo removes a todo item from the list
func (t *todoList) RemoveTodo(todoId int) {
	t.loadFromJson()

	for i, v := range t.todoStore {
		if v.ID == todoId {
			t.todoStore = append(t.todoStore[:i], t.todoStore[i+1:]...)
		}
	}

	t.saveToJson()
	fmt.Println("Removed todo with Id: ", todoId)
}

// ListTodo lists all todo items
func (t *todoList) ListTodo() {
	t.loadFromJson()

	for _, v := range t.todoStore {
		fmt.Printf("%d. %s - %t \n", v.ID, v.Task, v.Completed)
	}
}

// RemoveAllTodo removes all todo items
func (t *todoList) RemoveAllTodo() {
	t.loadFromJson()

	t.todoStore = []todo{}

	t.saveToJson()
	fmt.Println("Removed all todo items")
}

// CompleteTodo - This function marks a todo item as completed
func (t *todoList) CompleteTodo(id int) {
	t.loadFromJson()

	for i, v := range t.todoStore {
		if v.ID == id {
			t.todoStore[i].Completed = true
		}
	}

	t.saveToJson()
	fmt.Println("Completed todo with Id: ", id)
}
