package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

// SaveToJson - This function saves the todo list to a json file
func (t todoList) saveToJson() {
	// Marshal the data from t.todoStore
	data, err := json.Marshal(t.todoStore)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("db.json", data, 0o644); err != nil {
		panic(err)
	}
}

// LoadFromJson - This function loads the todo list from a json file
func (t *todoList) loadFromJson() {
	// Create a file if it doesn't exist
	if _, err := os.Stat("db.json"); os.IsNotExist(err) {
		_, err := os.Create("db.json")
		if err != nil {
			panic(err)
		}
	}

	data, err := os.ReadFile("db.json")
	if err != nil {
		panic(err)
	}

	if len(data) > 0 {
		// Unmarshal the data to t.todoStore
		err = json.Unmarshal(data, &t.todoStore)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("No todo's found")
	}
}

// newTodoId - This function returns a new todo id
func (t todoList) newTodoId() int {
	if len(t.todoStore) == 0 {
		return 1
	}
	return t.todoStore[len(t.todoStore)-1].ID + 1
}
