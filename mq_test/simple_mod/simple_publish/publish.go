package main

import (
	f "fmt"
	"github.com/Quaestiox/JustTalk_backend/rabbitmq"
)

func main() {
	rabbitmq := rabbitmq.NewRabbitMQSimple("justtalk")
	rabbitmq.PublishSimple("hello!")
	f.Println("sended successfully.")
}
