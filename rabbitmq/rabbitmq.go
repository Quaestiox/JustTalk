package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	Mqurl     string
}

func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	url := viper.GetString("MQ_URL")
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: url}
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

func (r *RabbitMQ) GetOneMsg() ([]byte, bool) {
	msg, ok, err := r.channel.Get(
		r.QueueName,
		true,
	)
	if err != nil {
		log.Print("get Message failed:", err)
		return nil, false
	}

	if !ok {
		log.Print("no message")
		return nil, false
	}

	return msg.Body, true
}
