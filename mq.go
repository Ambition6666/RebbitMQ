package mymq

import (
	"fmt"
	"log"

	mq "github.com/rabbitmq/amqp091-go"
)

var Url string

type MyRabbitMQ struct {
	conn    *mq.Connection
	channel *mq.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	Key string
	//连接信息
	Mqurl string
}

func InitMqConfig(username string, pwd string, host string, port string) {
	Url = fmt.Sprintf("amqp://%s:%s@%s:%s/", username, pwd, host, port)
}

// 创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *MyRabbitMQ {
	return &MyRabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: Url}
}

// 错误处理函数
func (r *MyRabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// 断开channel 和 connection
func (r *MyRabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}
