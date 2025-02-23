package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	"time"
)

const MQURL = "amqp://Quaestio:123567@:5672/master"

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	Mqurl     string
}

func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

func Unwrap(err error, message string) {
	if err != nil {
		log.Fatal().Err(err).Msg(message)
	}
}

func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queueName, "", "")
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	Unwrap(err, "failed to connect rabbitmq.")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	Unwrap(err, "failed to open channel.")
	return rabbitmq
}

func (r *RabbitMQ) PublishSimple(message string) {
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil)
	Unwrap(err, "")
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

func (r *RabbitMQ) ConsumeSimple() {
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil)
	Unwrap(err, "")
	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	Unwrap(err, "")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received one message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func (r *RabbitMQ) GetMsg() <-chan []byte {
	msgChan := make(chan []byte)
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil)
	Unwrap(err, "")

	go func() {

		for {
			// 获取队列的消息数量
			qInfo, err := r.channel.QueueInspect(q.Name)
			if err != nil {
				log.Print("队列检查失败:", err)
				return
			}

			if qInfo.Messages == 0 {
				// 通知 RabbitMQ 取消消费，退出循环
				err := r.channel.Cancel("consumer_tag", false)
				if err != nil {
					log.Print("取消消费者失败:", err)
				}
				return
			}
			time.Sleep(1 * time.Second) // 每秒检查一次
		}
	}()
	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	Unwrap(err, "")
	go func() {
		for d := range msgs {
			msgChan <- d.Body
		}
	}()

	return msgChan
}
func (r *RabbitMQ) GetOneMsg() ([]byte, bool) {
	msg, ok, err := r.channel.Get(
		r.QueueName,
		true, // autoAck = true, 自动确认消息
	)
	if err != nil {
		log.Print("获取消息失败:", err)
		return nil, false
	}

	if !ok {
		log.Print("队列为空，没有新消息")
		return nil, false
	}

	return msg.Body, true
}
