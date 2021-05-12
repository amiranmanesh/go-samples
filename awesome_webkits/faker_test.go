package main_test

import (
	"awesome_webkits/fake"
	"testing"
)

func TestFaker(t *testing.T) {

	result, _ := fake.Call("ipv6")
	t.Log(result)
}
