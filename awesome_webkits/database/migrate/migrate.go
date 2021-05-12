package migrate

import (
	"awesome_webkits/database/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/juju/errors"
	"gorm.io/gorm"
)

func DBMigrate(db *gorm.DB) *gorm.DB {

	if err :=
		db.AutoMigrate(
			&models.User{},
			&models.OauthAccessToken{},
			&models.Project{},
			&models.OauthProjectToken{},
			&models.ProjectApi{},
			&models.Email{},
		); err != nil {
		panic(errors.Trace(err))
	}

	return db
}
