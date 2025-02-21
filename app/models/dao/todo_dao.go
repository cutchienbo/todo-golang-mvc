package dao

import (
	"errors"
	"golang-mvc/app/helpers"
	"golang-mvc/app/models/db"
	"golang-mvc/app/models/requests"
	"golang-mvc/app/models/responses"
	"math"
	"time"
)

func FilterTodoExec(userId uint, req *requests.FilterTodoResquest, res *responses.FilterTodoResponse) error {
	var status string
	var dateFrom string = req.DateFrom
	var dateTo string = req.DateTo
	var priority []int
	
	if req.Status == "All" {
		status = "%%"
	} else {
		status = "%" + req.Status + "%"
	}

	if dateFrom == "" {
		dateFrom = time.Date(0000, time.January, 1, 0, 0, 0, 0, time.UTC).String()
	}

	if dateTo == "" {
		dateTo = time.Date(9999, time.December, 31, 23, 59, 59, 0, time.UTC).String()
	}

	if req.Priority == 0 {
		priority = []int{1, 2, 3}
	} else {
		priority = []int{req.Priority}
	}

	var totalRecord int64

	query := helpers.GormDB.Debug().Table("todo").
		Where("title LIKE ?", "%"+req.Title+"%").
		Where("status LIKE ?", status).
		Where("due_date BETWEEN ? AND ?", dateFrom, dateTo).
		Where("priority IN (?)", priority).
		Where("user_id = ?", userId).
		Where("deleted_at IS NULL")

	if err := query.Count(&totalRecord).Error; err != nil {
		return err
	}

	if totalRecord == 0 {
		return errors.New("todolist empty")
	}

	query = query.Offset((req.CurrentPage - 1) * req.Limit).Limit(req.Limit)

	if err := query.Find(&res.Todos).Error; err != nil {
		return err
	}

	res.TotalPage = int(math.Ceil(float64(totalRecord) / float64(req.Limit)))

	return nil
}

func CreateTodoExec(userId uint, req *requests.CreateTodoRequest, res *responses.CreateTodoResponse) error {
	var todo db.Todo = db.Todo{
		Title: req.Title,
		Description: req.Description,
		Status: "Pending",
		DueDate: req.DueDate,
		Priority: req.Priority,
		UserId: int64(userId),
	}

	if err := helpers.GormDB.Create(&todo); err == nil {
		return errors.New("cannot create new todo")
	}

	res.Todo = todo

	return nil
}

func UpdateTodoExec(req *requests.UpdateTodoRequest) error {
	var todo db.Todo

	helpers.GormDB.Debug().First(&todo, req.TodoId)
	helpers.GormDB.Debug().Model(&todo).Update(req.UpdateField, req.UpdateValue)
	helpers.GormDB.Debug().Model(&todo).Update("updated_at", helpers.GetCurrentTimeVN().String())

	return nil
}

func DeleteTodoExec(req *requests.DeleteTodoRequest) error {
	var todo db.Todo

	helpers.GormDB.Debug().First(&todo, req.TodoId)
	helpers.GormDB.Debug().Model(&todo).Update("deleted_at", helpers.GetCurrentTimeVN().String())

	return nil
}