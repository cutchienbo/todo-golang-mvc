package app

import (
	"golang-mvc/app/configs"
	"golang-mvc/app/db"
	"golang-mvc/app/helpers"
	"golang-mvc/app/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine{
	gin.SetMode(os.Getenv("GIN_MODE"))

	helpers.GormDB = configs.GormConnection()

	db.Migrate(helpers.GormDB)
	db.SeedDB(helpers.GormDB)

	r := routes.InitRoute()

	return r
}