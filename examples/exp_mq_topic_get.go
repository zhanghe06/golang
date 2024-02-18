package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)

func failOnErrorGet(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func connectionCloseGet(conn *amqp.Connection) {
	_ = conn.Close()
}

func channelCloseGet(ch *amqp.Channel) {
	_ = ch.Close()
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnErrorGet(err, "Failed to connect to RabbitMQ")
	defer connectionCloseGet(conn)

	ch, err := conn.Channel()
	failOnErrorGet(err, "Failed to open a channel")
	defer channelCloseGet(ch)

	ex := "project.test.topic"
	bk := []string{"shanghai.*", "*.huangpu"}
	mq := "test"
	//ct := "ctag-1"

	err = ch.ExchangeDeclare(
		ex,                 // name
		amqp.ExchangeTopic, // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)
	failOnErrorGet(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		mq,    // name if empty, Generate random name（eg：amq.gen-Zm2WyvxMbl1iqalEokvVsg）
		true,  // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnErrorGet(err, "Failed to declare a queue")

	for _, rk := range bk {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			q.Name, ex, rk)
		err = ch.QueueBind(
			q.Name, // queue name
			rk,     // routing key
			ex,     // exchange
			false,
			nil)
		failOnErrorGet(err, "Failed to bind a queue")
	}

	// 限制预取数量 [https://www.rabbitmq.com/confirms.html#channel-qos-prefetch]
	err = ch.Qos(1, 0, true)
	failOnErrorGet(err, "Failed to set channel qos")

	msg, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnErrorGet(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msg {
			log.Printf(" [x] %s, %d", d.Body, d.DeliveryTag)

			seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
			numRand := seededRand.Intn(3)
			// 随机场景
			switch numRand {
			// 确认删除
			case 0:
				time.Sleep(10 * time.Second)
				log.Printf(" [x] Ack")
				_ = d.Ack(false)
			// 重新入队
			case 1:
				log.Printf(" [x] Nack")
				// 如果不确认，则一直等待，当进程退出时，消息才会返回队列
				_ = d.Nack(false, true)
			// 立即丢弃
			case 2:
				log.Printf(" [x] Reject")
				_ = d.Reject(false)
			}
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
