package main

import (
	"github.com/Quaestiox/JustTalk_backend/rabbitmq"
)

func main() {
	rabbitmq := rabbitmq.NewRabbitMQSimple("justtalk")
	rabbitmq.ConsumeSimple()
}
