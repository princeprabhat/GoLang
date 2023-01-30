package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	database "github.com/princeprabhat/ToDoApi/Database"
	models "github.com/princeprabhat/ToDoApi/Models"
)

func main() {

	db := database.InitDb()

	newTodo := models.NewTodo(db)

	route := gin.Default()

	route.GET("/todosall", newTodo.GetTodo)
	route.GET("/todobyid/:id", newTodo.GetTodoById)
	route.POST("/createTodo", newTodo.CreateTodo)
	route.PATCH("/updateTodo/:id", newTodo.UpdateTodo)
	route.DELETE("/deteleTodo/:id", newTodo.DeleteBlog)

	route.Run()
	defer db.Close()

}

// connect to sqllite database
