package db

import "gorm.io/gorm"

type Todo struct {
	gorm.Model

	Title       string 	`gorm:"size:255;not null" json:"title"`        
    Description string 	`gorm:"type:text" json:"description"`          
    Status      string 	`gorm:"size:20;default:'pending'" json:"status"` 
    DueDate     string 	`gorm:"type:date" json:"due_date"`             
    Priority    int    	`gorm:"default:2" json:"priority"`   
	UserId		int64	`gorm:"default:null" json:"user_id"`
	User		User 	`gorm:"foreignKey:UserId;references:ID;constraint:OnDelete:CASCADE"`
}