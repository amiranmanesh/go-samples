package main

func generator(max int) <-chan int {
	outChInt := make(chan int, 100)

	go func() {
		for i := 0; i <= max; i++ {
			outChInt <- i
		}
		close(outChInt)
	}()

	return outChInt
}

func power(in <-chan int) <-chan int {
	outChInt := make(chan int, 100)

	go func() {
		for v := range in {
			outChInt <- v * v
		}
		close(outChInt)
	}()

	return outChInt
}

func sum(in <-chan int) <-chan int {
	outChInt := make(chan int, 100)

	go func() {
		var sum int
		for v := range in {
			sum += v
		}
		outChInt <- sum
		close(outChInt)
	}()

	return outChInt
}

func main() {
	amount := 10

	result := <-sum(power(generator(amount)))

	println(result)
}
