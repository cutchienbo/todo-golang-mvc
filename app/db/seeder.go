package db

import (
	"golang-mvc/app/helpers"
	"golang-mvc/app/models/db"
	"strconv"

	"gorm.io/gorm"
)

func SeedDB(gormDB *gorm.DB) {
	var count int64

	gormDB.Debug().Model(&db.User{}).Count(&count)

	if count > 0 {
		println("Database already seeded!")

		return
	}

	println("Seeding database ...")

	var users []db.User

	for i := 0; i < 3; i++ {
		todos := []db.Todo{
			{
				Title: "Do laundry",
				Description: "Take a clothing basket to a washing machine and use two cup of laundry detergent.",
				Status: "Pending",
				DueDate: "17:00:00 17/02/2025",
				Priority: 2,
			},
			{
				Title: "Water vegetable",
				Description: "Water cabbage with manure.",
				Status: "Completed",
				DueDate: "6:00:00 17/02/2025",
				Priority: 3,
			},
		}

		var password string = "user" + strconv.Itoa(i);

		var hashPassword, err = helpers.HashPassword(password)

		if err != nil {
			println("Hash password error!")

			return
		}

		users = append(users, db.User{
			Name: "User-" + strconv.Itoa(i),
			Password: hashPassword,
			Status: true,
			Todos: todos,
		})
	}

	gormDB.Create(&users)
}