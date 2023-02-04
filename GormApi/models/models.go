package models

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/princeprabhat/GormApi/utils"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string `json:code`
	Price uint   `json:price`
}

type ProductSotre struct {
	Products []Product
	Db       *gorm.DB
}

func StartProduct(database *gorm.DB) *ProductSotre {
	return &ProductSotre{
		Products: []Product{},
		Db:       database,
	}
}

func (p ProductSotre) GetAllProducts(c *gin.Context) {
	var product []Product
	p.Db.AutoMigrate(&Product{})

	p.Db.Find(&product)

	utils.Showresult(c, http.StatusOK, product, "All product fetched successfully")

}

func (p ProductSotre) CreateProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		utils.Showresult(c, http.StatusBadRequest, nil, "Error creating Todo")
		return
	}
	p.Db.AutoMigrate(&Product{})

	p.Db.Create(&Product{Code: product.Code, Price: product.Price})

	utils.Showresult(c, http.StatusOK, nil, "Product Created Successfully")

}

func (p ProductSotre) GetOneProduct(c *gin.Context) {
	var product Product
	str := c.Param("id")
	id, err := strconv.Atoi(str)
	utils.Checkerror(err)

	p.Db.AutoMigrate(&Product{})
	p.Db.First(&product, id)

	utils.Showresult(c, http.StatusOK, product, "Product fetched successfully")
}

func (p ProductSotre) UpdateProduct(c *gin.Context) {
	var product Product

	if err := c.ShouldBindJSON(&product); err != nil {
		utils.Showresult(c, http.StatusBadRequest, nil, "Error creating Todo")
		return
	}
	str := c.Param("id")
	id, err := strconv.Atoi(str)
	utils.Checkerror(err)

	p.Db.AutoMigrate(&Product{})

	p.Db.Model(&Product{}).Where("ID = ?", id).Updates(Product{Code: product.Code, Price: product.Price})

	data := fmt.Sprintf("Product with ID %d has been updated", id)
	utils.Showresult(c, http.StatusOK, data, "Product Updated Successfully")

}

func (t ProductSotre) DeleteProduct(c *gin.Context) {
	str := c.Param("id")
	id, err := strconv.Atoi(str)
	utils.Checkerror(err)
	t.Db.Delete(&Product{}, id)
	data := fmt.Sprintf("Product id %d has been deleted", id)

	utils.Showresult(c, http.StatusOK, data, "Product Deleted Successfully")
}
