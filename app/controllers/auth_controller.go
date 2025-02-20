package controllers

import (
	"golang-mvc/app/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Check Valid Token
// @Description Use to check jwt token
// @Tags auth
// @Accept application/json
// @Produce json
// @Router /auth/checkToken [get]
func CheckValidJWT(c *gin.Context) {
	var jwt string = c.GetHeader("Authorization")

	if err := helpers.CheckJWT(jwt); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
				"message": "token invalid",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "token valid",
		},
	)
}