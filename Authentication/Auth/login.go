package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	middleware "github.com/princeprabhat/Authentication/Middleware"
	models "github.com/princeprabhat/Authentication/Models"
	utils "github.com/princeprabhat/Authentication/Utils"
	"golang.org/x/crypto/bcrypt"
)

func (a AuthController) Login(c *gin.Context) {
	var user models.User
	var loginuser models.Login
	if err := c.ShouldBindJSON(&loginuser); err != nil {
		utils.Showresult(c, http.StatusBadRequest, nil, "Error getting User")
		return
	}

	userInst := a.Db.First(&user, "email = ?", loginuser.Email)
	if userInst.RowsAffected == 0 {
		utils.Showresult(c, http.StatusNotFound, nil, "User not found")
		return
	}

	if loginuser.Email != user.Email {
		utils.Showresult(c, http.StatusNotFound, nil, "Wrong Credentials")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginuser.Password)); err != nil {
		utils.Showresult(c, http.StatusNotFound, nil, "Wrong Password")
		return
	}

	if user.IsActive != true {
		utils.Showresult(c, http.StatusNotImplemented, user.IsActive, "user is not active, please verify the user")
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &middleware.Claims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(middleware.JwtKey)

	if err != nil {
		utils.Showresult(c, http.StatusInternalServerError, nil, "Error generating token")
		return
	}

	email := claims.Email

	var m map[string]any = map[string]any{
		"email": email,
		"token": tokenString,
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	utils.Showresult(c, http.StatusOK, m, "User Fetched successfully!")

}

func (a AuthController) VerifyUser(c *gin.Context) {
	var Verify models.Login
	var usermodel models.User

	if err := c.BindJSON(&Verify); err != nil {
		utils.Showresult(c, http.StatusNoContent, nil, err.Error())
		return
	}

	userInst := a.Db.First(&usermodel, "email = ?", Verify.Email)
	if userInst.RowsAffected == 0 {
		utils.Showresult(c, http.StatusNotFound, nil, "User not found")
		return
	}

	if Verify.Email != usermodel.Email {
		utils.Showresult(c, http.StatusNotFound, nil, "Wrong Email")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usermodel.Password), []byte(Verify.Password)); err != nil {
		utils.Showresult(c, http.StatusNotFound, nil, "Wrong Password")
		return
	}

	if usermodel.IsActive == true {
		utils.Showresult(c, http.StatusForbidden, nil, "User is already verified")
		return
	} else {
		usermodel.IsActive = true
	}

	a.Db.Save(&usermodel)

	utils.Showresult(c, http.StatusOK, nil, "User Verified Successfully")

}
