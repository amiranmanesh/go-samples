package lib

import (
	"awesome_webkits/utils/random"
	"awesome_webkits/utils/txt"
	"fmt"
)

const (
	defaultColorsCount = 1
)

func (p *Person) setColors() {
	var colors []string
	var count int
	if p.Favorites.ColorsCount == 0 {
		count = defaultActivitiesCount
	} else {
		count = p.Favorites.ColorsCount
	}
	for i := 0; i < count; i++ {
		colors = append(colors, RandomColor())
	}
	p.Favorites.ColorsCount = count
	p.Favorites.Colors = colors
}

func RandomColor() string {

	colors := txt.GetFileData("./fake/db/colors.txt")
	colorsSize := len(colors)
	return colors[random.RandInt(0, colorsSize)]
}

func RandomRGBColor() string {
	r := random.RandInt(0, 255)
	g := random.RandInt(0, 255)
	b := random.RandInt(0, 255)

	return fmt.Sprintf("%d,%d,%d", r, g, b)
}
