package routes

import (
	"golang-mvc/app/controllers"
	"golang-mvc/app/middlewares"
	"net/http"

	// "golang-mvc/app/middlewares"
	_ "golang-mvc/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoute() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		// Thêm header CORS cho mỗi request
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Cho phép tất cả các origin
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS") // Các phương thức HTTP cho phép
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Các header cho phép

		// Xử lý preflight OPTIONS
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		
		c.Next()
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//version 1.0
	v1 := r.Group("/api/v1")
	{
		//auth
		v1.POST("/login", controllers.UserLoginController)
		v1.POST("/register", controllers.UserRegisterController)
		v1.GET("/auth/checkToken", controllers.CheckValidJWT)

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