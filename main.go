package main

import (
	AuthController "inkxk/jwt-api/controller/auth"
	UserController "inkxk/jwt-api/controller/user"
	"inkxk/jwt-api/middleware"
	"inkxk/jwt-api/orm"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// connect database
	orm.ConnectDb()

	// router
	r := gin.Default()
	r.Use(cors.Default())

	// route
	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)

	authorized := r.Group("/user", middleware.Token())
	authorized.GET("/get-all", UserController.GetAllUsers)
	authorized.GET("/get-profile", UserController.GetUserProfile)

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
