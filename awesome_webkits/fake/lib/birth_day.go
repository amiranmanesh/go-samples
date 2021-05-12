package lib

import (
	"awesome_webkits/utils/random"
	"time"
)

func (p *Person) setBirthDay(minAge, maxAge int) {
	p.BirthDate = RandomBirthDay()
}

func RandomBirthDay() time.Time {

	minAge := 10
	maxAge := 50

	year := time.Now().Year() - maxAge + random.RandInt(0, maxAge-minAge)
	month := random.RandInt(1, 13)
	day := random.RandInt(1, 29)
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func (p Person) getAge() int {
	year, _, _, _, _, _ := timeDiff(p.BirthDate, time.Now())
	return year
}

func timeDiff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}
