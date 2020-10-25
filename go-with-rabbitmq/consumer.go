package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://localhost:5672")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to rabbitmq...")
	defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
    	log.Fatal(err)
	}
	defer ch.Close()

    messages, err := ch.Consume(
    	"BusyBoy",
    	"",
    	true,
    	false,
    	false,
    	false,
    	nil,
    	)

    forever := make(chan bool)
    go func() {
    	for d := range messages {
			fmt.Printf("Received message is %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever
}