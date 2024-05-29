package user

import (
	"fmt"
	"inkxk/jwt-api/model"
	"inkxk/jwt-api/orm"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var responseBody model.GetAllUserResponse
	var users []orm.User

	authorization := c.Request.Header.Get("Authorization")
	tokenString := strings.Replace(authorization, "Bearer ", "", 1)
	fmt.Println(tokenString)
	// token, err := jwt.Parse(tokenStr, nil)
	// if token == nil {
	// 	return nil, err
	// }
	// claims, _ := token.Claims.(jwt.MapClaims)

	// query users
	orm.Db.Find(&users)

	responseBody.Status = http.StatusOK
	responseBody.Message = "registerd success"
	responseBody.Users = users

	c.JSON(responseBody.Status, responseBody)
}
