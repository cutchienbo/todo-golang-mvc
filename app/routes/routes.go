package routes

import (
	"golang-mvc/app/controllers"
	"golang-mvc/app/middlewares"
	// "golang-mvc/app/middlewares"
	_ "golang-mvc/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoute() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//version 1.0
	v1 := r.Group("/api/v1")
	{
		//auth
		v1.POST("/login", controllers.UserLoginController)
		v1.POST("/register", controllers.UserRegisterController)

		//todo
		v1.POST("/", middlewares.AuthGuard, controllers.GetFilterTodosController)
		v1.POST("/create", middlewares.AuthGuard, controllers.CreateTodoController)
		v1.PATCH("/update", middlewares.AuthGuard, controllers.UpdateTodoController)
		v1.DELETE("/delete", middlewares.AuthGuard, controllers.DeleteTodoController)
	}

	//test
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}