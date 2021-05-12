package lib

import "time"

type Person struct {
	Title             string         `json:"title"`
	FirstName         string         `json:"first_name"`
	MiddleName        string         `json:"middle_name"`
	LastName          string         `json:"last_name"`
	Username          string         `json:"username"`
	NationalCode      string         `json:"national_code"`
	NationalCodeCount int            `json:"-"`
	Gender            string         `json:"gender"`
	BirthDate         time.Time      `json:"birth_date"`
	Email             string         `json:"email"`
	AvatarUrl         string         `json:"avatar_url"`
	Favorites         PersonFavorite `json:"favorites"`
}

type PersonFavorite struct {
	Activities      []string `json:"favorite_activities"`
	ActivitiesCount int      `json:"-"`
	Colors          []string `json:"favorite_colors"`
	ColorsCount     int      `json:"-"`
	Foods           []string `json:"favorite_foods"`
	FoodsCount      int      `json:"-"`
	Fruits          []string `json:"favorite_fruits"`
	FruitsCount     int      `json:"-"`
}

func PersonGenerator() Person {
	p := Person{}
	p.setGender()
	p.setTitle()
	p.setName()
	p.setBirthDay(18, 50)
	p.setUsername()
	p.setEmail()
	p.setAvatar()
	p.setNationalCode()
	p.setActivities()
	p.setColors()
	p.setFoods()
	p.setFruits()
	return p
}
