package src

import (
	"github.com/streadway/amqp"
	"log"
)

func Connect() *amqp.Channel {

	log.Println("Starting rabbitmq connection...")
	conn, err := amqp.Dial("amqp://localhost:5672")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to rabbitmq...")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
    log.Println("Channel creation is OK")

	return ch

}
