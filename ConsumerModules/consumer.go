package main

import (
	"PublisherSubscriberAssignment/ConsumerModules/models"
	"PublisherSubscriberAssignment/ConsumerModules/services"
	"encoding/json"
	"fmt"
	"log"

	"PublisherSubscriberAssignment/ConsumerModules/helpers/confighelpers"
	"PublisherSubscriberAssignment/ConsumerModules/helpers/loghelpers"

	"github.com/streadway/amqp"
)

// main function
func main() {

	// initialize configs from current directory
	confighelpers.InitViper(".")

	// Create connection
	conn, err := amqp.Dial(confighelpers.GetConfig("consumerURL"))
	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		failOnError(err, "Failed to open a channel")
	}
	defer ch.Close()

	// declare a queue
	q, err := DeclareQueue(ch)
	if err != nil {
		failOnError(err, "Failed to declare a queue")
	}

	// consume a message from the queue
	msgs, err := ConsumeQueue(&q, ch)
	if err != nil {
		failOnError(err, "Failed to Consume a message")
	}

	neverending := make(chan bool)

	go func() {
		// looping through the messages
		for d := range msgs {
			// transmission object declaration
			transmittedObject := &models.IncomingMessage{}
			err := json.Unmarshal(d.Body, transmittedObject)
			if err != nil {
				fmt.Println("Error..!!", err)
			} else {
				loghelpers.Info("Received a message: ", transmittedObject)
				// call to save service for saving the data
				status := services.SaveHotelData(transmittedObject)
				if status == nil {
					loghelpers.Info("Data Saved Successfully")
				} else {
					loghelpers.Error("Failed to save data")
				}

			}
		}
	}()

	loghelpers.Info(" [*] Waiting for messages. To exit press CTRL+C")
	<-neverending
}

// DeclareQueue declares and returns new queue object
func DeclareQueue(ch *amqp.Channel) (queue amqp.Queue, err error) {

	queue, err = ch.QueueDeclare(
		"hotel_assignment", // name
		false,              // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	return
}

// ConsumeQueue declares and return delivary of messages
func ConsumeQueue(queue *amqp.Queue, ch *amqp.Channel) (msgs <-chan amqp.Delivery, err error) {

	msgs, err = ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	return
}

// error logs
func failOnError(err error, msg string) {
	log.Fatalf("%s: %s", msg, err)
}
