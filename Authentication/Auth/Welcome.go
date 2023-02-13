package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	middleware "github.com/princeprabhat/Authentication/Middleware"
	models "github.com/princeprabhat/Authentication/Models"
	utils "github.com/princeprabhat/Authentication/Utils"
)

func (a AuthController) Welcome(c *gin.Context) {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			utils.Showresult(c, http.StatusUnauthorized, nil, "Please Login again")
			return
		}
		utils.Showresult(c, http.StatusInternalServerError, nil, "Internal Server Error")
		return

	}
	tknstr := cookie.Value

	claims := &middleware.Claims{}

	tkn, err := jwt.ParseWithClaims(tknstr, claims, func(token *jwt.Token) (interface{}, error) {
		return middleware.JwtKey, nil

	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			utils.Showresult(c, http.StatusUnauthorized, nil, "Unauthorize, invalid signature")
			return
		}
		utils.Showresult(c, http.StatusBadRequest, nil, "Bad Request")
		return
	}
	if !tkn.Valid {
		utils.Showresult(c, http.StatusUnauthorized, nil, "Token is invalid")
		return
	}
	var userprofile models.User

	a.Db.First(&userprofile, "email = ?", claims.Email)

	str := fmt.Sprintf("Welcome %s", userprofile.FirstName)

	utils.Showresult(c, http.StatusOK, userprofile, str)

}

func (a AuthController) Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})

	utils.Showresult(c, http.StatusOK, nil, "User logged out successfully")
}
