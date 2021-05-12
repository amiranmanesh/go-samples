package lib

import (
	"awesome_webkits/utils/random"
	"fmt"
	"strings"
)

func (p *Person) setUsername() {

	switch rr := random.RandInt(0, 6); rr {
	case 0:
		p.Username = strings.ToLower(fmt.Sprintf("%s.%s", p.FirstName, p.LastName))
	case 1:
		p.Username = strings.ToLower(fmt.Sprintf("%c%s", p.FirstName[0], p.LastName))
	case 2:
		p.Username = strings.ToLower(fmt.Sprintf("%s-%s", p.LastName, p.FirstName))
	case 3:
		p.Username = strings.ToLower(fmt.Sprintf("%c.%c.%s", p.FirstName[0], p.MiddleName[0], p.LastName))
	case 4:
		p.Username = strings.ToLower(fmt.Sprintf("%s%d", p.FirstName, p.BirthDate.Year()))
	case 5:
		p.Username = strings.ToLower(fmt.Sprintf("%s%s%d", p.FirstName, p.LastName, p.BirthDate.Year()))
	}
}

func RandomUsername() string {

	return fmt.Sprintf("%s_%s", RandomSurName(), RandomSurName())
}
