package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte("ORXWWZLOON2HE2LOM5ZWKY3VOJSXAYLTON3W64TE")

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func VerifyJwtToken(c *gin.Context) string {
	tokenStr := c.GetHeader("Token")

	log.Println(tokenStr)
	return tokenStr
}
