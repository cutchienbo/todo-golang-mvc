package dao

import (
	"errors"
	"golang-mvc/app/helpers"
	"golang-mvc/app/models/db"
	"golang-mvc/app/models/requests"
)

func UserLoginExec(req *requests.UserLoginRequest, userJWTSubject *helpers.UserJWTSubject) (error) {
	var user *db.User

	if res := helpers.GormDB.Debug().Table("user").Where("name = ?", req.Name).Find(&user); res.RowsAffected == 0 {
		return errors.New("username wrong")
	}

	if res := helpers.CheckPasswordHash(req.Password, user.Password); !res {
		return errors.New("password wrong")
	}	

	userJWTSubject.Id = user.ID
	userJWTSubject.Name = user.Name

	return nil
}

func UserRegisterExec(req *requests.UserRegisterRequest, userJWTSubject *helpers.UserJWTSubject) (error) {
	var user *db.User

	if res := helpers.GormDB.Debug().Table("user").Where("name = ?", req.Name).Find(&user); res.RowsAffected > 0 {
		return errors.New("username existed")
	}

	if req.Password != req.RePassword {
		return errors.New("confirm password wrong")
	}

	if hashPassword, err := helpers.HashPassword(req.Password); err == nil {
		user.Name = req.Name
		user.Password = hashPassword

		helpers.GormDB.Create(&user)

		return nil
	}

	return errors.New("cannot hash password")
}