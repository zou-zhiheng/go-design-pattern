package simpleFactory

import (
	"fmt"
)

//简单工厂模式
//go语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类
//NewXXX函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐就是简单工厂
//在这个simpleFactory包中只有API接口和NewAPI函数为包外可见，封装了实现的细节

//API是interface类型
type API interface {
	Say(name string) string
}

//创建API时会返回一个API类型实例
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

//hiAPI是API的一种实现
type hiAPI struct{}

//Say hi to name
func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi,%s", name)
}

//helloAPI是API的另外一种实现
type helloAPI struct{}

//say hello to name
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello,%s", name)
}
