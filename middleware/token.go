package middleware

import (
	"fmt"
	"inkxk/jwt-api/model"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		var responseBody model.MainResponse

		// get bearer token
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			responseBody.Status = http.StatusUnauthorized
			responseBody.Message = "invalid token"
			c.AbortWithStatusJSON(responseBody.Status, responseBody)
		}

		tokenString := strings.Replace(authorization, "Bearer ", "", 1)

		// validate token
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
		} else {
			responseBody.Status = http.StatusUnauthorized
			responseBody.Message = "invalid token"
			c.AbortWithStatusJSON(responseBody.Status, responseBody)
		}

		c.Next()
	}
}
