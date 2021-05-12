package lib

import (
	"awesome_webkits/utils/random"
	"fmt"
	"time"
)

func RandomTimestamp(minYear, maxYear int) int64 {
	if maxYear == 0 {
		maxYear = time.Now().Year()
	}

	year := random.RandInt(minYear-maxYear, time.Now().Year()-maxYear)
	month := random.RandInt(1, 11)
	day := random.RandInt(1, 29)

	now := time.Now()
	return now.AddDate(year, month, day).Unix()

}

func RandomDate(minYear, maxYear int, seperator string) string {
	if maxYear == 0 {
		maxYear = time.Now().Year()
	}

	year := random.RandInt(minYear, maxYear)
	month := random.RandInt(1, 11)
	day := random.RandInt(1, 29)

	return fmt.Sprintf("%d%s%d%s%d", year, seperator, month, seperator, day)

}
