package main

import "sync"

type Counter struct {
	sync.Mutex
	value int
}

func main() {
	counter := Counter{}
	var wait sync.WaitGroup
	wait.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			counter.Lock()
			counter.value++
			defer counter.Unlock()
			defer wait.Done()
		}(i)
	}
	wait.Wait()
	println(counter.value)
}
