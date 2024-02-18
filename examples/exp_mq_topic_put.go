package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func failOnErrorPut(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func connectionClosePut(conn *amqp.Connection) {
	_ = conn.Close()
}

func channelClosePut(ch *amqp.Channel) {
	_ = ch.Close()
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnErrorPut(err, "Failed to connect to RabbitMQ")
	defer connectionClosePut(conn)

	ch, err := conn.Channel()
	failOnErrorPut(err, "Failed to open a channel")
	defer channelClosePut(ch)

	notifyClose := make(chan *amqp.Error)
	notifyConfirm := make(chan amqp.Confirmation)
	ch.NotifyClose(notifyClose)     // channel connection 断开均通知
	ch.NotifyPublish(notifyConfirm) // 消息发送确认通知

	// 通道设置为Confirm模式
	err = ch.Confirm(false)
	failOnErrorPut(err, "confirm.select destination")

	ex := "project.test.topic"
	//rk := "shanghai.huangpu"
	rk := "guangzhou.huangpu"
	//rk := "guangzhou.liwan"

	// 测试gin项目
	ex = "ex.project.topic"
	rk = "rk.project"

	err = ch.ExchangeDeclare(
		ex,                 // name
		amqp.ExchangeTopic, // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // no-wait
		nil,                // arguments
	)
	failOnErrorPut(err, "Failed to declare an exchange")

	body := "test msg"
	err = ch.Publish(
		ex,    // exchange
		rk,    // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(body),
			DeliveryMode: amqp.Persistent, // Transient (0 or 1) or Persistent (2)
		})
	failOnErrorPut(err, "Failed to publish a message")

	// channel 测试关闭（此时，关闭，消息有可能已经发送成功，只是没有确认成功）
	//err = ch.Close()
	//if err != nil {
	//	fmt.Printf(err.Error())
	//}
	// 确认超时设置
	confirmTimeout := time.After(time.Duration(5) * time.Second)

	select {
	case confirm := <-notifyConfirm:
		if !confirm.Ack {
			//log.Printf(" [x] ERROR: publish confirm error")
			err = fmt.Errorf("%v", "publish confirm error")
		}
	case <-notifyClose:
		err = fmt.Errorf("%v", "channel close")
	case <-confirmTimeout:
		err = fmt.Errorf("%v", "confirm timeout")
	}
	failOnErrorPut(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s, ex: %s, rk: %s", body, ex, rk)
}
