package user

import (
	"fmt"
	"inkxk/jwt-api/model"
	"inkxk/jwt-api/orm"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetAllUsers(c *gin.Context) {
	var responseBody model.GetAllUserResponse
	var users []orm.User

	// get bearer token
	authorization := c.Request.Header.Get("Authorization")
	tokenString := strings.Replace(authorization, "Bearer ", "", 1)

	// validate token
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	// validate claim
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// query users
		orm.Db.Find(&users)

		// success response
		responseBody.Status = http.StatusOK
		responseBody.Message = "registerd success"
		responseBody.Users = users
	} else {
		responseBody.Status = http.StatusUnauthorized
		responseBody.Message = "invalid token"
	}

	// if !token.Valid {
	// 	responseBody.Status = http.StatusInternalServerError
	// 	responseBody.Message = "jwt parse token error"
	// 	c.JSON(responseBody.Status, responseBody)
	// 	return
	// }

	c.JSON(responseBody.Status, responseBody)
}
