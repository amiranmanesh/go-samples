package lib

import "awesome_webkits/utils/random"

var (
	genderMale   = "male"
	genderFemale = "female"
	gender       = [...]string{genderMale, genderFemale}
)

func (p *Person) setGender() {
	p.Gender = RandomGender()
}

func RandomGender() string {
	return gender[random.RandInt(0, len(gender))]
}
