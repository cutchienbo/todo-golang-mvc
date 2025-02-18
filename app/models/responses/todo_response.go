package responses

import "golang-mvc/app/models/db"

type FilterTodoResponse struct {
	Todos []db.Todo
}

type CreateTodoResponse struct {
	Todo db.Todo
}

type UpdateTodoResponse struct {
	UpdateField string
	UpdateValue string
}