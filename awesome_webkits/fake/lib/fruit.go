package lib

import (
	"awesome_webkits/utils/random"
	"awesome_webkits/utils/txt"
)

const (
	defaultFruitsCount = 1
)

func (p *Person) setFruits() {
	var fruits []string
	var count int
	if p.Favorites.FruitsCount == 0 {
		count = defaultFruitsCount
	} else {
		count = p.Favorites.FruitsCount
	}
	for i := 0; i < count; i++ {
		fruits = append(fruits, RandomFruit())
	}
	p.Favorites.FruitsCount = count
	p.Favorites.Fruits = fruits
}

func RandomFruit() string {
	fruits := txt.GetFileData("./fake/db/fruits.txt")
	fruitsSize := len(fruits)
	index := random.RandInt(0, fruitsSize)
	return fruits[index]
}
