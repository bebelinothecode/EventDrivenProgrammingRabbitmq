package main

import (
	"log"
	"github.com/bebelino/internal"
)


func main() {
	conn, err := internal.ConnectRabbitMQ("bebelino","secret","localhost:5672","customers")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client, error := internal.NewRabbitMQClient(conn)

	if error != nil {
		panic(error)
	}

	messageBus, err := client.ConsumeMessage("customers_created","email-service",false)

	if err != nil {
		panic(err)
	}

	var blocking chan struct{}

	go func() {
		for message := range messageBus {
			log.Printf("Received a message:%s",message.Body)

			if err := message.Ack(false); err != nil {
				log.Println("Acknowledge message failed")
				continue
			}
			log.Printf("Acknowledge message:%s\n",message.MessageId)
		}
	}()

	log.Println("Consuming,to close the program press CTRL+C")
	<-blocking
}