package random

import (
	guuid "github.com/google/uuid"
	"github.com/rs/xid"
	"math/rand"
	"time"
)

var (
	randInitialized bool = false
)

//c01d7cf6-ec3f-47f0-9556-a5d6e9009a43
func RandString36Bytes() string {
	id := guuid.New()
	return id.String()
}

//b50vl5e54p1000fo3gh0
func RandString20Bytes() string {
	id := xid.New()
	return id.String()
}

func RandInt(min int, max int) int {
	if !randInitialized {
		randInit()
	}
	return min + rand.Intn(max-min)
}

func randInit() {
	rand.Seed(time.Now().UTC().UnixNano())
	randInitialized = true
}
