package main

import (
	"context"
	"log"
	"time"
	"github.com/bebelino/internal"
	"github.com/rabbitmq/amqp091-go"
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

	err = client.CreateQueue("customers_created", true, false)

	if err != nil {
		panic(err)
	}

	err = client.CreateQueue("customers_test", false, true)

	if err != nil {
		panic(err)
	}

	err = client.CreateBinding("customers_created", "customers.created.*", "customers")

	if err != nil {
		panic(err)
	}

	err = client.CreateBinding("customers_test","customers.*","customers")

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Minute)

	defer cancel()

	err = client.SendMessage(ctx, "customers", "customers.created.us", amqp091.Publishing{
		ContentType: "text/plain",
		DeliveryMode: amqp091.Persistent,
		Body: []byte("A cool message between services"),
	})

	if err != nil {
		panic(err)
	}

	err = client.SendMessage(ctx, "customers", "customers.test", amqp091.Publishing{
		ContentType: "text/plain",
		DeliveryMode: amqp091.Transient,
		Body: []byte("Bebelino is transient"),
	})

	if err != nil {
		panic(err)
	}

	time.Sleep(3 * time.Minute)

	log.Println(&client)
}