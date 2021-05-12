package utils

import (
	"awesome_webkits/database"
	"fmt"
)

type iDbUtils interface {
	IsExists(table string, filed string, value string) bool
}

type dbUtils struct{}

var DbUtils iDbUtils = &dbUtils{}

func (dbUtils) IsExists(table string, field string, value string) bool {

	var result int

	query := fmt.Sprintf("SELECT count(*) FROM  %s where %s = '%s'", table, field, value)
	fmt.Println(query)
	database.DB.GetDB().Raw(query).Scan(&result)
	fmt.Println(result)
	return result > 0
}
