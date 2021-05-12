package lib

import "awesome_webkits/utils/random"

var (
	maleTitles   = [...]string{"Mr"}
	femaleTitles = [...]string{"Mrs", "Miss", "Ms"}
	uniSexTitles = [...]string{"Dr", "Prof", "Rev"}
)

func (p *Person) setTitle() {
	if p.Gender == genderMale {
		if random.RandInt(0, 2) == 0 {
			p.setMaleTitle()
		} else {
			p.setUniSexTitle()
		}
	} else if p.Gender == genderFemale {
		if random.RandInt(0, 2) == 0 {
			p.setFemaleTitle()
		} else {
			p.setUniSexTitle()
		}
	} else {
		p.setUniSexTitle()
	}
}

func (p *Person) setMaleTitle() {
	p.Title = RandomMaleTitle()
}

func (p *Person) setFemaleTitle() {
	p.Title = RandomFemaleTitle()
}

func (p *Person) setUniSexTitle() {
	p.Title = RandomUniSexTitle()
}

func RandomMaleTitle() string {
	return maleTitles[random.RandInt(0, len(maleTitles))]
}

func RandomFemaleTitle() string {
	return femaleTitles[random.RandInt(0, len(femaleTitles))]
}

func RandomUniSexTitle() string {
	return uniSexTitles[random.RandInt(0, len(uniSexTitles))]
}
