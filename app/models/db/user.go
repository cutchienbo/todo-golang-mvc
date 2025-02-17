package db

import "gorm.io/gorm"

type User struct {
	gorm.Model 

	Name 		string	`gorm:"not null;size:64;uniqueIndex" json:"name"`
	Password 	string	`gorm:"not null;size:255" json:"password"`
	Status		bool	`gorm:"default:true" json:"status"`
	Todos 		[]Todo
}