package lib

import (
	"awesome_webkits/utils/random"
	"fmt"
	"strings"
)

func RandomIpV4() string {

	values := make([]interface{}, 4, 4)
	for i := range values {
		values[i] = random.RandInt(1, 254)
	}
	return fmt.Sprintf("%d.%d.%d.%d", values...)
}

func RandomIpV6() string {
	var attributes = []string{
		"0", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e",
	}

	values := make([]interface{}, 8, 8)
	blocks := make([]string, 4, 4)
	for i := range values {

		for index := range blocks {
			blocks[index] = attributes[random.RandInt(0, len(attributes))]
		}

		values[i] = strings.Join(blocks, "")
	}

	return fmt.Sprintf("%s:%s:%s:%s:%s:%s:%s:%s", values...)
}
