package unit_test

import (
	"fmt"
	"iris-demo-new/entity/request"
	"iris-demo-new/services"
	"strconv"
	"testing"
)

// OrderCh 双向管道的初始化，可放在程序的入口处，初始化1次，即可
var OrderCh chan request.OrderInfo
var submitNumber = 100
func init()  {
	OrderCh = make(chan request.OrderInfo, submitNumber)
	recOrder := services.NewRecOrder()
	// 先消费, for...select...case 语句 阻塞 获取单向接受管道的数据
	go recOrder.HandleOrder(OrderCh)
}

func TestSubmitOrder(t *testing.T)  {
	sendOrder := services.NewSendOrder()
	// 模拟n次并发请求
	for i:=0; i<submitNumber; i++ {
		str := strconv.Itoa(i)
		params := request.OrderInfo{
			Id: uint64(i),
			Sn: "JH20220328172" + str,
		}
		// 再在任何地方生产
		sendOrder.SendOrder(OrderCh, params)
	}
}

func TestChannel(t *testing.T)  {
	c := make(chan int) //   chan   //读写

	go counter(c) //生产者
	printer(c)    //消费者

	fmt.Println("done")
}

//   chan<- //只写
func counter(out chan<- int) {
	defer close(out)
	for i := 0; i < 5; i++ {
		fmt.Println("i=",i)
		out <- i //如果对方不读 会阻塞
	}
}

//   <-chan //只读
func printer(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}



// 生产者消费者模型实现
type OrderInfo struct {
	Id uint64
	Sn string
}

// send
type sendIter interface {
	send(out chan<- OrderInfo)
}

type send struct {

}

func (s *send) send(out chan<- OrderInfo, params OrderInfo) {
	fmt.Println("request to send Id=", params.Id)
	out <- OrderInfo{
		Id: params.Id,
		Sn: params.Sn,
	}
}

func SendOrder(out chan<- OrderInfo, params OrderInfo)  {
	sendOrder := new(send)
	sendOrder.send(out, params)
}


// receive
type recIter interface {
	rec(in <-chan OrderInfo)
}

type rec struct {

}

func (s *rec) rec(in <-chan OrderInfo) {
	select {
	case item := <-in:
		// save to db or other operations...
		fmt.Printf("Id=%d, Sn=%s \n", item.Id, item.Sn)
	}
	//for item := range in {
	//	// save to db or other operations...
	//	fmt.Printf("Id=%d, Sn=%s \n", item.Id, item.Sn)
	//}
}

func HandleOrder(in <-chan OrderInfo)  {
	recOrder := new(rec)
	recOrder.rec(in)
}

func TestOrder(t *testing.T)  {
	n := 100
	ch := make(chan OrderInfo)
	defer close(ch)
	// 模拟n次并发请求
	for i:=0; i<n; i++ {
		str := strconv.Itoa(i)
		params := OrderInfo{
			Id: uint64(i),
			Sn: "JH20220318172" + str,
		}
		go SendOrder(ch, params)
		HandleOrder(ch)
	}
}




