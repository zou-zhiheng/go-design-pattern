package RabbitMQ

func recieve() {
	rabbitmq := NewRabbitMQSimple("" + "sttch")
	rabbitmq.ConsumeSimple()
}
