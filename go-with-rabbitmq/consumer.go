package main

import (
	"fmt"
	"go-small-projects/go-with-rabbitmq/src"
	"log"
)

func main() {

    messages, err := src.Connect().Consume(
    	"BusyBoy",
    	"",
    	true,
    	false,
    	false,
    	false,
    	nil,
    	)
    if err != nil {
    	log.Fatal(err)
	}

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