package main

import (
	"awesome_webkits/cache"
	"awesome_webkits/database"
	"awesome_webkits/database/migrate"
	"awesome_webkits/http"
	"awesome_webkits/utils/validation"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.WithFields(logrus.Fields{
		"package": "main",
		"func":    "main",
	}).Info("Start ...")

	database.DB.Initialize()
	migrate.DBMigrate(database.DB.GetDB())
	validation.SetCustomValidations()
	cache.Cache.InitRedisClientInstance()
	http.App.Initialize()
	http.App.Run()
}
