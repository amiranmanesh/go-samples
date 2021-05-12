package lib

import (
	"awesome_webkits/utils/random"
	"fmt"
)

var (
	emailYahoo   = "yahoo.com"
	emailGmail   = "gmail.com"
	emailHotmail = "hotmail.com"
	emailGmx     = "gmx.com"
	emailsHost   = [...]string{emailYahoo, emailGmail, emailHotmail, emailGmx}
)

func (p *Person) setEmail() {
	p.Email = RandomEmail()
}

func RandomEmail() string {
	username := RandomUsername()

	return fmt.Sprintf("%s@%s", username, emailsHost[random.RandInt(0, len(emailsHost))])
}
