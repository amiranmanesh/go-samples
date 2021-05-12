package random

import (
	"testing"
)

func TestRandom(t *testing.T) {
	min := 10000
	max := -10000
	for i := 0; i < 10000; i++ {
		random := RandInt(20, 50)
		if random < min {
			min = random
		}
		if random > max {
			max = random
		}
	}
	t.Logf("min :%d , max :%d", min, max)
}
