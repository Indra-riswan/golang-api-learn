package handler

import (
	"learn2/helper"
	"learn2/service"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJwt(service service.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authheader := c.GetHeader("Authorization")
		if authheader == "" {
			respons := helper.BuildErrorRespons("Failed to sent request", "no token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, respons)
			return
		}
		token, err := service.ValidateToken(authheader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[user_id] :", claims["user_id"])
			log.Println("Claims[issuer] :", claims["issuer"])
		} else {
			log.Println(err)
			respons := helper.BuildErrorRespons("Token Invalid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
		}
	}
}
