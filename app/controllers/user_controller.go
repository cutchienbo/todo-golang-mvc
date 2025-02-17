package controllers

import (
	"golang-mvc/app/helpers"
	"golang-mvc/app/models/dao"
	"golang-mvc/app/models/requests"
	"golang-mvc/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Login 
// @Description Login todo website by name and password
// @Tags users
// @Accept json
// @Produce json 
// @Param name query string true "User name"
// @Param password query string true "User password"
// @Success 200 {object} UserLoginResponse "Login successfully"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Router /api/v1/login [get]
func UserLoginController(c *gin.Context) {
	var req requests.UserLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"errors": gin.H{
					"message": err.Error(),
				},
			},
		)

		return
	}

	var userJWTSubject helpers.UserJWTSubject

	if err := dao.UserLoginExec(&req, &userJWTSubject); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"errors": gin.H{
					"message": err.Error(),
				},
			},
		)

		return
	}

	var response responses.UserLoginResponse

	response.Token = helpers.GenerateToken(userJWTSubject)

	c.JSON(
		http.StatusOK, gin.H{
			"message": "Login successfully",
			"data": response,
		},
	)
}

func UserRegisterController(c *gin.Context) {
	
}