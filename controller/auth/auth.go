package auth

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"inkxk/jwt-api/model"
	"inkxk/jwt-api/orm"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var requestBody RegisterBody
	var responseBody model.RegisterResponse
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check user exists
	var userExist orm.User
	orm.Db.Where("username = ?", requestBody.Username).First(&userExist)
	if userExist.ID > 0 {
		responseBody.Status = http.StatusConflict
		responseBody.Message = "user existed"
		c.JSON(responseBody.Status, responseBody)
		return
	}

	// encrypt password
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(requestBody.Password), 10)

	// insert database
	user := orm.User{Username: requestBody.Username, Email: requestBody.Email, Password: string(encryptedPassword)}
	result := orm.Db.Create(&user)
	fmt.Printf("registered: %d rows. \n", result.RowsAffected)

	// reponse
	if user.ID > 0 {
		responseBody.Status = http.StatusOK
		responseBody.Message = "registerd success"
		responseBody.UserId = user.ID
	} else {
		responseBody.Message = "registered failed"
	}

	c.JSON(responseBody.Status, responseBody)
}

func Login(c *gin.Context) {
	var requestBody LoginBody
	var responseBody model.LoginResponse
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check user exists
	var user orm.User
	orm.Db.Where("username = ?", requestBody.Username).First(&user)
	if user.ID == 0 {
		responseBody.Status = http.StatusNotFound
		responseBody.Message = "user not found"
		c.JSON(responseBody.Status, responseBody)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password))
	if err == nil {
		// gen access key
		claims := &jwt.MapClaims{
			"iss": "issuer",
			"exp": time.Now().Add(time.Hour * 24).Unix(),
			"data": map[string]string{
				"user_id": strconv.FormatUint(uint64(user.ID), 10),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString(os.Getenv("JWT_SECRET_KEY"))

		responseBody.Status = http.StatusOK
		responseBody.Message = "login success"
		responseBody.AccessToken = tokenString
	} else {
		responseBody.Status = http.StatusUnauthorized
		responseBody.Message = "invalid password"
	}

	c.JSON(responseBody.Status, responseBody)
}
