package main

import (
	AuthController "inkxk/jwt-api/controller/auth"
	"inkxk/jwt-api/orm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// connect database
	orm.ConnectDb()

	// router
	r := gin.Default()
	r.Use(cors.Default())

	// route
	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
