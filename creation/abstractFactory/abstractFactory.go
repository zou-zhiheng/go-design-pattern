package abstractFactory

import "fmt"

//抽象工厂模式
//用于生产产品族的工厂，所生成的对象是有关联的
//如果抽象工厂退化成的对象无关联则成为工厂函数模式
//比如本例子中使用的RDB和XML存储订单信息,抽象工厂分别能生成相关的主订单信息和订单详情信息
//如果业务逻辑中需要替换使用的时候只需要改动工厂函数相关的类就能替换使用不同的存储方式了


//OrderMainDAO 订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

//OrderDetailDAO 订单详情记录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

//DAOFactory DAO 抽象模式工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

//RDBMainDAO 关系型数据库的OrderMainDAO实现
type RDBMainDAO struct {}

//SaveOrderMain 实现
func(*RDBMainDAO) SaveOrderMain(){
	fmt.Println("rdb main save")
}

//RDBDetailDAO 关系型数据库OrderDetailDAO实现
type RDBDetailDAO struct {}

//SaveOrderDetail 实现
func (*RDBDetailDAO) SaveOrderDetail()  {
	fmt.Println("rdb detail save")
}

//RDBDAOFactory 是RDB 抽象工厂实现
type RDBDAOFactory struct {}



//创建订单主记录
func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

//创建订单详情记录
func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO  {
	return &RDBDetailDAO{}
}

//XMLMainDAO XML存储
type XMLMainDAO struct{}

//SaveOrderMain ...
func (*XMLMainDAO) SaveOrderMain() {
	fmt.Println("xml main save")
}

//XMLDetailDAO XML存储
type XMLDetailDAO struct{}

// SaveOrderDetail ...
func (*XMLDetailDAO) SaveOrderDetail() {
	fmt.Println("xml detail save")
}

//XMLDAOFactory 是RDB 抽象工厂实现
type XMLDAOFactory struct{}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}
