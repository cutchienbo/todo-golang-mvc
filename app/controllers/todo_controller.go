package controllers

import (
	"golang-mvc/app/helpers"
	"golang-mvc/app/models/dao"
	"golang-mvc/app/models/requests"
	"golang-mvc/app/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Filter Todo
// @Description Use to filter todo by title, status, due date, priority
// @Tags todos
// @Accept application/json
// @Produce json
// @Param filterData body string true "Todo title, status, dateFrom, dateStart, priority"
// @Router / [post]
func GetFilterTodosController(c *gin.Context) {
	var userId uint = helpers.GetTokenSubject(c.GetHeader("Authorization")).Id

	var req requests.FilterTodoResquest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
				"message": "cannot bind params",
			},
		)

		return
	}

	var res responses.FilterTodoResponse

	if err := dao.FilterTodoExec(userId, &req, &res); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
				"message": "cannot find todo",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "get todo successfully",
			"data": res,
		},
	)
}

// @Summary Create Todo
// @Description Use to create todo
// @Tags todos
// @Accept application/json
// @Produce json
// @Param createData body string true "Todo title, description, dueDate, priority"
// @Router /create [post]
func CreateTodoController(c *gin.Context) {
	var userId uint = helpers.GetTokenSubject(c.GetHeader("Authorization")).Id

	var req requests.CreateTodoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
				"message": "cannot bind params",
			},
		)

		return
	}

	var res responses.CreateTodoResponse

	if err := dao.CreateTodoExec(userId, &req, &res); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
				"message": "cannot create new todo",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "create todo successfully",
			"data": res,
		},
	)
}

// @Summary Update Todo
// @Description Use to update todo
// @Tags todos
// @Accept application/json
// @Produce json
// @Param updateData body string true "Todo todoId, updateField, updateValue"
// @Router /update [patch]
func UpdateTodoController(c *gin.Context) {
	var req requests.UpdateTodoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
				"message": "cannot bind params",
			},
		)

		return
	}

	var res responses.UpdateTodoResponse

	if err := dao.UpdateTodoExec(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
				"message": "cannot update todo",
			},
		)

		return
	}

	res.UpdateField = req.UpdateField
	res.UpdateValue = req.UpdateValue

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "update todo successfully",
			"data": res,
		},
	)
}

// @Summary Delete Todo
// @Description Use to delete todo
// @Tags todos
// @Accept application/json
// @Produce json
// @Param deleteData body string true "Todo todoId"
// @Router /delete [delete]
func DeleteTodoController(c *gin.Context) {
	var req requests.DeleteTodoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
				"message": "cannot bind params",
			},
		)

		return
	}

	if err := dao.DeleteTodoExec(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
				"message": "cannot delete todo",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "delete todo successfully",
		},
	)
}