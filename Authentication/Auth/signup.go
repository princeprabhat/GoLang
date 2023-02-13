package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/princeprabhat/Authentication/Models"
	utils "github.com/princeprabhat/Authentication/Utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	Db *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		Db: db,
	}
}

//signup user from API server

func (a AuthController) SignUp(c *gin.Context) {
	var usersignup models.SignUpSchema
	var usermodel models.User
	if err := c.ShouldBindJSON(&usersignup); err != nil {
		utils.Showresult(c, http.StatusBadRequest, nil, "Error creating User")
		return
	}
	userIns := a.Db.Where("email = ?", usersignup.Email).First(&usermodel)
	if userIns.RowsAffected != 0 {
		utils.Showresult(c, http.StatusConflict, nil, "User already exists")
		return
	}
	if len(usersignup.Password) < 6 {
		utils.Showresult(c, http.StatusLengthRequired, nil, "Password must be at least 6 characters long")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usersignup.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.Showresult(c, http.StatusBadRequest, nil, "Error generating password")
		return
	}

	usermodel.Password = string(hashedPassword)
	usermodel.Email = usersignup.Email
	usermodel.FirstName = usersignup.FirstName
	usermodel.LastName = usersignup.LastName

	// Create the user
	if err := a.Db.Create(&usermodel).Error; err != nil {
		utils.Showresult(c, http.StatusInternalServerError, nil, "Error creating user")
		return
	}

	utils.Showresult(c, http.StatusOK, nil, "User Signup Successfull")

}
