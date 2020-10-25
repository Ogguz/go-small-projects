package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"go-small-projects/go-with-rabbitmq/src"
	"log"
	"time"
)

func main() {

	src.Connect()
	defer src.Connect().Close()

	// declare the queue
	que, err := src.Connect().QueueDeclare("BusyBoy",false,true,false,true,nil)
	fmt.Println(que)
	if err != nil {
		log.Panic(err)
	}
	log.Println("queue declared successfully")

    // Oh be finally publish them :)
    err = src.Connect().Publish(
    	"",
    	"BusyBoy",
    	true,
    	false,
    	amqp.Publishing{
			Headers:         nil,
			ContentType:     "text/plain",
			Timestamp:       time.Time{},
			Body:            []byte("Are you cola?"),
		},
	)
    if err != nil {
    	log.Panic(err)
	}

	log.Println("Successfully published message to the queue")
    fmt.Println("ok")
}
