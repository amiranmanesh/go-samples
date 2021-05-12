package database

import (
	"awesome_webkits/utils/env"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/juju/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type iDatabase interface {
	Initialize()
	GetDB() *gorm.DB
}

type database struct{}

var DB iDatabase = &database{}

var (
	db *gorm.DB
)

func (database) GetDB() *gorm.DB {
	if db == nil {
		db = connect()
	}
	return db
}

func (database) Initialize() {
	if db == nil {
		db = connect()
	}
}

func connect() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		env.GetEnvItem("DATABASE_USER"),
		env.GetEnvItem("DATABASE_PASSWORD"),
		env.GetEnvItem("DATABASE_HOST"),
		env.GetEnvItem("DATABASE_PORT"),
		env.GetEnvItem("DATABASE_NAME"),
	)
	connection, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic(errors.Trace(err))
	}
	return connection
}
