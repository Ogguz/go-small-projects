package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func main() {

    log.Println("Starting rabbitmq connection...")

    // connect to rabbitmq broker
    conn, err := amqp.Dial("amqp://localhost:5672") // buraya test de yaz
    if err != nil {
    	log.Fatal(err)
	}
	log.Println("Successfully connected to rabbitmq")
	defer conn.Close()  // TODO hem ch hem conn'a gerek var mi? sor

	// open up a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()
	log.Println("Channel created successfully")

	// declare the queue
	que, err := ch.QueueDeclare("BusyBoy",false,true,false,true,nil)
	fmt.Println(que)
	if err != nil {
		log.Panic(err)
	}
	log.Println("queue declared successfully")

    // Oh be finally publish them :)
    err = ch.Publish(
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
