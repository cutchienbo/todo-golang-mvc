package middlewares

import (
	"golang-mvc/app/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthGuard(c *gin.Context) {
	var jwt string = c.GetHeader("Authorization")

	if err := helpers.CheckJWT(jwt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"message": err.Error(),
				"status": http.StatusBadRequest,
			},
		})

		c.Abort()

		return
	}	
	
	c.Next()
}