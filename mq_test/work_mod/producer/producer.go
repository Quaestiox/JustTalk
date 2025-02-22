package main

import (
	f "fmt"
	"github.com/Quaestiox/JustTalk_backend/rabbitmq"
	"strconv"
	"time"
)

func main() {
	rabbitmq := rabbitmq.NewRabbitMQSimple("justtalk")

	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple(strconv.Itoa(i) + ":hello!")
		time.Sleep(1 * time.Second)
		f.Printf("%s:sended successfully.\n", i)
	}

}
