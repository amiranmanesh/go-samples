package main

import "time"

func main() {
	//sample 1
	/*
		channel := make (chan string)

		go func(){
			channel <- "hello world"
			println("finishing go routine")
		}()

		message := <- channel
		fmt.Println(message)
	*/

	//sample 2
	/*
		channel := make (chan string)
		var waitGroup sync.WaitGroup
		waitGroup.Add(1)
		go func(){
			channel <- "hello world"
			println("finishing go routine")
			waitGroup.Done()
		}()

		message := <- channel
		fmt.Println(message)
		waitGroup.Wait()
	*/

	//sample 3
	/*
		channel := make(chan string, 1)

		go func() {
			channel <- "hello world 1"
			channel <- "hello world 2"
			println("finishing go routine")
		}()

		message := <-channel
		fmt.Println(message)
	*/

	//sample 4
	/*
		channel := make(chan string, 2)

		go func(ch chan<- string) {
			//you cannot listen to this channel
			ch <- "hello world 1"
			ch <- "hello world 2"
			println("finishing go routine")
		}(channel)

		message1 := <-channel
		message2 := <-channel
		fmt.Println(message1)
		fmt.Println(message2)
	*/

	//sample 5
	/*
		helloCh := make(chan string, 1)
		goodByeCh := make(chan string, 1)
		quitCh := make(chan bool)
		go receiver(helloCh, goodByeCh, quitCh)
		go sendString(helloCh, "Hello world")
		time.Sleep(time.Second)
		go sendString(goodByeCh, "GoodBye")
		<-quitCh
	*/

	//sample 6
	ch := make(chan int)
	go func() {
		ch <- 1
		time.Sleep(time.Second)
		ch <- 2
		close(ch)
	}()
	for v := range ch {
		println(v)
	}
}

func sendString(ch chan<- string, s string) {
	ch <- s
}

func receiver(
	helloCh <-chan string,
	goodByeCh <-chan string,
	quitCh chan<- bool /*you cannot listen to this channel*/,
) {
	for {
		select {
		case msg := <-helloCh:
			println(msg)
		case msg := <-goodByeCh:
			println(msg)
		case <-time.After(time.Second * 2):
			println("Noting received in 2 seconds. Exiting...")
			quitCh <- true
			break
		}
	}
}
