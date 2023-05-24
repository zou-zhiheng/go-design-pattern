package RabbitMQ

import "fmt"

func publish(){
	rabbitmq:=NewRabbitMQSimple(""+"sttch")
	rabbitmq.PublishSimple("hello sttch2023")
	fmt.Println("发送成功")
}
