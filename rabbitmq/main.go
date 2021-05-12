package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

//docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
//http://localhost:15672/
func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	fmt.Println("Successfully connected to our RabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("Test Queue", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)

	for i := 0; i < 1000; i++ {

		err = ch.Publish("", "Test Queue", false, false, amqp.Publishing{
			ContentType: "test/plain",
			Body:        []byte("Hello world!"),
		})
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		fmt.Println("Successfully published Message to Queue")
		time.Sleep(time.Second)
	}

	fmt.Println(q)
}
