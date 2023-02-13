package models

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	utils "github.com/princeprabhat/Authentication/Utils"
	"gorm.io/gorm"
)

func StartUser(dbIns *gorm.DB) *NewUser {
	return &NewUser{
		Users: []User{},
		Db:    dbIns,
	}
}

func (n NewUser) GetUsers(c *gin.Context) {
	var users []User

	n.Db.AutoMigrate(&User{})

	n.Db.Find(&users)

	utils.Showresult(c, http.StatusOK, users, "All Users Fetched Successfully")

}

func (n NewUser) CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Showresult(c, http.StatusBadRequest, nil, "Error creating User")
		return
	}
	n.Db.AutoMigrate(&User{})

	n.Db.Create(&User{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Password: user.Password})

	utils.Showresult(c, http.StatusOK, nil, "User Created Successfully")

}

func (n NewUser) GetOneUser(c *gin.Context) {
	var user User
	str := c.Param("id")
	id, err := strconv.Atoi(str)
	utils.CheckError(err)

	n.Db.AutoMigrate(&User{})
	n.Db.First(&user, id)
	if id == int(user.ID) {
		utils.Showresult(c, http.StatusOK, user, "User fetched successfully")

	} else {
		utils.Showresult(c, http.StatusNotFound, nil, "User not found")
	}

}

func (n NewUser) UpdateUser(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Showresult(c, http.StatusBadRequest, nil, "Error Updating User")
		return
	}
	str := c.Param("id")
	id, err := strconv.Atoi(str)
	utils.CheckError(err)

	n.Db.AutoMigrate(&User{})

	n.Db.Model(&User{}).Where("ID = ?", id).Updates(User{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Password: user.Password})

	data := fmt.Sprintf("User number %d has been updated", id)
	utils.Showresult(c, http.StatusOK, data, "User Updated Successfully")

}

func (t NewUser) DeleteUser(c *gin.Context) {
	str := c.Param("id")
	id, err := strconv.Atoi(str)
	utils.CheckError(err)
	d1 := t.Db.Delete(&User{}, id)
	log.Println(d1)
	data := fmt.Sprintf("User number %d has been deleted", id)

	utils.Showresult(c, http.StatusOK, data, "User Deleted Successfully")
}
