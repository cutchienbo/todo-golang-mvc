package requests

type FilterTodoResquest struct {
	Title       string `json:"title" binding:"omitempty"`
	Status      string `json:"status" binding:"oneof=Pending Progressing Completed Canceled All"`
	DateFrom    string `json:"dateFrom" binding:"omitempty"`
	DateTo      string `json:"dateTo" binding:"omitempty"`
	Priority    int    `json:"priority" binding:"oneof=0 1 2 3"`
	CurrentPage int	   `json:"currentPage" binding:"required"`
	Limit		int	   `json:"limit" binding:"required"`
	// Title    string `json:"title"`
	// Status   string `json:"status"`
	// DateFrom string `json:"date_from"`
	// DateTo   string `json:"date_to"`
	// Priority int    `json:"priority"`
}

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"omitempty"`
	DueDate     string `json:"dueDate" binding:"required"`
	Priority    int    `json:"priority" binding:"oneof=1 2 3"`
	UserId      int64  `json:"userId" binding:"omitempty"`
}

type UpdateTodoRequest struct {
	TodoId      uint   `json:"todoId" binding:"required"`
	UpdateField string `json:"updateField" binding:"required,oneof=title description status due_date priority"`
	UpdateValue string `json:"updateValue" binding:"required"`
}

type DeleteTodoRequest struct {
	TodoId uint `json:"todoId" binding:"required"`
}