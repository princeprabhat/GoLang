package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	utils "github.com/princeprabhat/ToDoApi/Utils"
)

type Todo struct {
	Id    int    `json:id`
	Title string `json:title`
	Body  string `json:body`
}

type TodoStore struct {
	Todos []Todo
	Db    *sql.DB
}

func NewTodo(databaseConn *sql.DB) *TodoStore {
	return &TodoStore{
		Todos: []Todo{},
		Db:    databaseConn,
	}
}

func (t TodoStore) GetTodo(c *gin.Context) {
	query, err := t.Db.Query("SELECT * FROM todo")
	utils.CheckError(err)

	defer query.Close()

	for query.Next() {
		var numTodo Todo
		err := query.Scan(&numTodo.Id, &numTodo.Title, &numTodo.Body)
		utils.CheckError(err)

		t.Todos = append(t.Todos, numTodo)

	}
	utils.ShowResult(c, http.StatusOK, t.Todos, "Printed Todos Succesfully")

}

func (t TodoStore) GetTodoById(c *gin.Context) {
	str := c.Param("id")
	id, err := strconv.Atoi(str)
	utils.CheckError(err)
	query, err := t.Db.Query("Select * from todo WHERE id = ?", id)
	utils.CheckError(err)

	defer query.Close()

	var resultTodo Todo
	for query.Next() {
		err := query.Scan(&resultTodo.Id, &resultTodo.Title, &resultTodo.Body)
		utils.CheckError(err)
		t.Todos = append(t.Todos, resultTodo)

	}
	if resultTodo.Id == id {
		utils.ShowResult(c, http.StatusOK, t.Todos, "Todo fetched successfully")

	} else {
		utils.ShowResult(c, http.StatusBadRequest, nil, "no todos found")
	}

}

func (t TodoStore) CreateTodo(c *gin.Context) {
	var todoList Todo
	if err := c.ShouldBindJSON(&todoList); err != nil {
		utils.ShowResult(c, http.StatusBadRequest, nil, "Error creating Todo")
		return
	}

	query, err := t.Db.Prepare("INSERT INTO todo(title,body) VALUES(?,?)")
	utils.CheckError(err)

	res, err := query.Exec(todoList.Title, todoList.Body)
	utils.CheckError(err)

	id, err := res.LastInsertId()
	utils.CheckError(err)

	utils.ShowResult(c, http.StatusOK, id, "Todo Created Succesfully")

}

func (t TodoStore) UpdateTodo(c *gin.Context) {
	var updateList Todo
	if err := c.ShouldBindJSON(&updateList); err != nil {
		utils.ShowResult(c, http.StatusBadRequest, nil, "Error updating Todo")
		return
	}

	str := c.Param("id")
	id, err := strconv.Atoi(str)
	utils.CheckError(err)

	query, err := t.Db.Query("UPDATE todo SET title =?,body = ? WHERE id = ?", updateList.Title, updateList.Body, id)

	defer query.Close()

	for query.Next() {
		err := query.Scan(&updateList.Id, &updateList.Title, &updateList.Body)
		utils.CheckError(err)

		t.Todos = append(t.Todos, updateList)

	}

	data := fmt.Sprintf("Todo number %d has been updated", id)
	utils.ShowResult(c, http.StatusOK, data, "Todo updated successfully")

}

func (t TodoStore) DeleteBlog(c *gin.Context) {
	str := c.Param("id")
	id, err := strconv.Atoi(str)
	utils.CheckError(err)

	query, err1 := t.Db.Prepare("DELETE FROM todo WHERE id=?")
	utils.CheckError(err1)

	_, err2 := query.Exec(id)
	utils.CheckError(err2)

	data := fmt.Sprintf("Todo number %d has been deleted", id)

	utils.ShowResult(c, http.StatusOK, data, "Todo deleted Successfully")

}
