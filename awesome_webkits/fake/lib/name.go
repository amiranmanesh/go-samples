package lib

import (
	"awesome_webkits/utils/random"
	"awesome_webkits/utils/txt"
)

func (p *Person) setName() {
	p.LastName = RandomSurName()
	p.MiddleName = RandomMaleName()
	if p.Gender == genderMale {
		p.FirstName = RandomMaleName()
	} else {
		p.FirstName = RandomFemaleName()
	}
}

func GetFullName(isMale bool) string {
	if isMale {
		return RandomMaleName() + " " + RandomSurName()
	} else {
		return RandomFemaleName() + " " + RandomSurName()
	}

}

func RandomMaleName() string {
	maleNames := txt.GetFileData("./fake/db/names_male.txt")
	namesSize := len(maleNames)
	return maleNames[random.RandInt(0, namesSize)]
}

func RandomFemaleName() string {
	femaleNames := txt.GetFileData("./fake/db/names_female.txt")
	namesSize := len(femaleNames)
	return femaleNames[random.RandInt(0, namesSize)]
}

func RandomSurName() string {
	surNames := txt.GetFileData("./fake/db/surnames.txt")
	namesSize := len(surNames)
	return surNames[random.RandInt(0, namesSize)]
}
