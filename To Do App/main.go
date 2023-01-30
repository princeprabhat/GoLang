package main

import (
	"fmt"
	"os"

	"github.com/rama-kairi/todo/cmd"
)

func main() {
	// td := todo.NewTodoService()

	// // go run main.go add Learn PostgreSQL
	// args := os.Args[1:]

	// switch args[0] {
	// case "add":
	// 	td.AddTodo(args[1:])
	// case "remove":
	// 	if len(args[1:]) > 1 {
	// 		fmt.Println("Invalid command")
	// 		return
	// 	}
	// 	id, err := strconv.Atoi(args[1])
	// 	if err != nil {
	// 		fmt.Println("Invalid command")
	// 		return
	// 	}
	// 	td.RemoveTodo(id)
	// case "list":
	// 	if len(args[1:]) > 0 {
	// 		fmt.Println("Invalid command")
	// 		return
	// 	}
	// 	td.ListTodo()
	// default:
	// 	fmt.Println("Invalid command")
	// }

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
