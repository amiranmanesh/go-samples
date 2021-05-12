package lib

import (
	"awesome_webkits/utils/random"
	"awesome_webkits/utils/txt"
)

const (
	defaultFoodsCount = 1
)

func (p *Person) setFoods() {
	var foods []string
	var count int
	if p.Favorites.FoodsCount == 0 {
		count = defaultFoodsCount
	} else {
		count = p.Favorites.FoodsCount
	}
	for i := 0; i < count; i++ {
		foods = append(foods, RandomFood())
	}
	p.Favorites.FoodsCount = count
	p.Favorites.Foods = foods
}

func RandomFood() string {

	foods := txt.GetFileData("./fake/db/foods.txt")
	foodsSize := len(foods)
	index := random.RandInt(0, foodsSize)

	return foods[index]
}
