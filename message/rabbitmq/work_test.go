// Author: yangzq80@gmail.com
// Date: 2021-05-27
//
package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"testing"
)

func BenchmarkSend(b *testing.B) {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	log.Println("Starting send...", b.N)

	b.ResetTimer()
	//for i := 0; i < b.N; i++ {
	for i := 0; i < 1000000; i++ {
		body := "Hello World!" + strconv.Itoa(i)
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Failed to publish a message")
		//log.Printf(" [x] Sent %s", body)
	}
}
