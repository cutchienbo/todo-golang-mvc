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
// @Accept application/json
// @Produce json
// @Param name&password body string true "User name and password"
// @Router /login [post]
func UserLoginController(c *gin.Context) {
	var req requests.UserLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"errors": err.Error(),
				"message": "cannot bind params",
			},
		)

		return
	}

	var userJWTSubject helpers.UserJWTSubject

	if err := dao.UserLoginExec(&req, &userJWTSubject); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"errors": err.Error(),
				"message": "login failed",
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

// @Summary Register 
// @Description Register todo website by name, password and rePassword
// @Tags users
// @Accept application/json
// @Produce json 
// @Param name&password&rePassword body string true "User name, password, rePassword"
// @Router /register [post]
func UserRegisterController(c *gin.Context) {
	var req requests.UserRegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"errors": err.Error(),
				"message": "cannot bind params",
			},
		)

		return
	}

	var userJWTSubject helpers.UserJWTSubject

	if err := dao.UserRegisterExec(&req, &userJWTSubject); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"errors": err.Error(),
				"message": "register failed",
			},
		)

		return
	}

	var response responses.UserLoginResponse

	response.Token = helpers.GenerateToken(userJWTSubject)

	c.JSON(
		http.StatusOK, gin.H{
			"message": "Register successfully",
			"data": response,
		},
	)
}