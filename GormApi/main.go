package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	database "github.com/princeprabhat/GormApi/Database"
	"github.com/princeprabhat/GormApi/models"
)

func main() {
	db := database.InitDatabase()

	newProduct := models.StartProduct(db)

	route := gin.Default()
	route.GET("/GetProducts", newProduct.GetAllProducts)
	route.POST("/CreateProduct", newProduct.CreateProduct)
	route.GET("/SingleProduct/:id", newProduct.GetOneProduct)
	route.PATCH("/UpdateProduct/:id", newProduct.UpdateProduct)
	route.DELETE("/DeleteProduct/:id", newProduct.DeleteProduct)
	route.Run()
}
