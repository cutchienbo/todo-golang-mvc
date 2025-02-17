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
		return errors.New("user name wrong")
	}

	if res := helpers.CheckPasswordHash(req.Password, user.Password); !res {
		return errors.New("user password wrong")
	}	

	userJWTSubject.Id = user.ID
	userJWTSubject.Name = user.Name

	return nil
}