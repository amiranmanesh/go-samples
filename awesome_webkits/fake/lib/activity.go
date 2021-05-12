package lib

import (
	"awesome_webkits/utils/txt"
)

const (
	defaultActivitiesCount = 3
)

func (p *Person) setActivities() {
	var activities []string
	var count int
	if p.Favorites.ActivitiesCount == 0 {
		count = defaultActivitiesCount
	} else {
		count = p.Favorites.ActivitiesCount
	}
	for i := 0; i < count; i++ {
		activities = append(activities, RandomActivity(i))
	}
	p.Favorites.ActivitiesCount = count
	p.Favorites.Activities = activities
}

func RandomActivity(index int) string {
	activities := txt.GetFileData("./fake/db/activities.txt")
	activitiesSize := len(activities)
	return activities[index%activitiesSize]
}
