package lib

import (
	"awesome_webkits/utils/random"
	"fmt"
)

const (
	defaultNationCount = 10
)

func (p *Person) setNationalCode() {
	p.NationalCode = RandomNationalCode()
}

func RandomNationalCode() string {
	var code string
	var count int
	count = defaultNationCount

	code += fmt.Sprintf("%d", random.RandInt(1, 10))
	for i := 1; i < count; i++ {
		code += fmt.Sprintf("%d", random.RandInt(0, 10))
	}
	return code
}
