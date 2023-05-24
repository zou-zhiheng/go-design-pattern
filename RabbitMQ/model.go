package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
)

const MQURL = "amqp://sttch:sttch@sttch@106.55.171.176:5672/sttch"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind key 名称
	Key string
	//连接信息
	Mqurl string
}

//创建结构体实例
func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: exchange, Mqurl: MQURL}
}

//断开channel和connection
func (r *RabbitMQ) Destroy() {
	err := r.channel.Close()
	if err != nil {
		fmt.Println(r.QueueName, "关闭失败")
		return
	}
	err = r.conn.Close()
	if err != nil {
		fmt.Println(r.conn.Major, "关闭失败")
	}
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		fmt.Println(message, err)
		panic(fmt.Sprintf(message, err))
	}
}

//创建简单模式下RabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	//创建RabbitMQ实例
	rabbitmq := NewRabbitMQ(queueName, "", "")
	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "failed to connect rabbitmq")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open channel")

	return rabbitmq
}

//直接模式队列生产
func (r *RabbitMQ) PublishSimple(message string) {
	//申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否具有排他性
		false,
		//是否自动删除
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	//调用channel 发送消息到队列中
	err=r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true,根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err!=nil{
		fmt.Println(r.QueueName,"发送失败")
	}
}

//simple 模式下消费者
func (r *RabbitMQ) ConsumeSimple() {
	//申请队列，如果队列不存在会自动创建,存在则跳过创建
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	//接收消息
	msg, err := r.channel.Consume(
		q.Name, //queue
		//区分消费者
		"", //consumer
		//是否自动应答
		true,
		//是否独有
		false,
		//设置为true，表示不能将同一个Connection中生产者发送的消息传递给这个Connection中的消费者
		false,
		//是否阻塞
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range msg {
			//消息逻辑处理，可以自行设计逻辑
			fmt.Println("Received a message:", string(d.Body))
		}
	}()

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
