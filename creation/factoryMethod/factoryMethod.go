package factoryMethod

//工厂方法模式
//工厂方法模式使用了子类的方式延迟生成对象到子类中实现
//Go中不存在继承，所以用匿名组合实现

//Operator是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

//OperatorFactory是工厂接口
type OperatorFactory interface {
	Create() Operator
}

//OperatorBase 是Operation接口实现的基类,封装公用方法
type OperatorBase struct {
	a,b int
}

//SetA 设置a
func (o *OperatorBase) SetA(a int){
	o.a=a
}

//SetB 设置b
func (o *OperatorBase) SetB(b int){
	o.b=b
}

//PlusOperatorFactory是PlusOperator的工厂类
type PlusOperatorFactory struct {}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

//PlusOperator Operator的实际方法实现
type PlusOperator struct {
	*OperatorBase
}

//Result获取结果
func (o PlusOperator) Result() int{
	return o.a+o.b
}

//MinusOperatorFactory是MinusOperator工厂类
type MinusOperatorFactory struct {}

func (MinusOperatorFactory) Create() Operator{
	return &MinusOperator{
		OperatorBase:&OperatorBase{},
	}
}

//MinusOperator Operator 实际减法实现
type MinusOperator struct {
	*OperatorBase
}

//Result 获取结果
func (o MinusOperator) Result() int{
	return o.a-o.b
}