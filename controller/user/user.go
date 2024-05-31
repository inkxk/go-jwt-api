package user

import (
	"inkxk/jwt-api/model"
	"inkxk/jwt-api/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var responseBody model.GetAllUserResponse
	var users []orm.User

	// query users
	orm.Db.Find(&users)

	// success response
	responseBody.Status = http.StatusOK
	responseBody.Message = "query success"
	responseBody.Users = users

	c.JSON(responseBody.Status, responseBody)
}

func GetUserProfile(c *gin.Context) {
	var responseBody model.GetUserProfileResponse
	var user orm.User

	userId := c.MustGet("user_id").(string)

	// query users
	orm.Db.First(&user, userId)

	// success response
	responseBody.Status = http.StatusOK
	responseBody.Message = "query success"
	responseBody.User = user

	c.JSON(responseBody.Status, responseBody)
}
