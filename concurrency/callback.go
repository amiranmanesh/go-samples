package main

import (
	"fmt"
	"strings"
	"sync"
	"syscall/js"
)

func main() {
	var wait sync.WaitGroup
	wait.Add(1)
	toUpperAsync("hello world", func(v string) {
		toUpperAsync(fmt.Sprintf("callback: %s\n",v),func(v string) {
			fmt.Printf("callback: %s\n",v)
			wait.Done()
		})
	})
	println("Waiting async response...")
	wait.Wait()
}

func toUpperAsync(word string, f func(string)) {
	go func() {
		f(strings.ToUpper(word))
	}()
}
