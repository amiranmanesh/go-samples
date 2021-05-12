package lib

import (
	"fmt"
	"time"
)

func (p *Person) setAvatar() {
	p.AvatarUrl = RandomAvatar(&p.Username)
}

func RandomAvatar(uniqueIndex *string) string {
	if uniqueIndex == nil {
		return fmt.Sprintf("https://i.pravatar.cc/250?u=%d", time.Now().Unix())
	} else {
		return fmt.Sprintf("https://i.pravatar.cc/250?u=%s", *uniqueIndex)
	}
}
