package services

import (
	"fmt"
	"iris-demo-new/entity/request"
)

// send
type sendIter interface {
	send(out chan<- request.OrderInfo, params request.OrderInfo)
}

type send struct {

}

func NewSendOrder() *send {
	return &send{}
}

func (s *send) send(out chan<- request.OrderInfo, params request.OrderInfo) {
	fmt.Println("request to send Id=", params.Id)
	out <- request.OrderInfo{
		Id: params.Id,
		Sn: params.Sn,
	}
}

func (s *send) SendOrder(out chan<- request.OrderInfo, params request.OrderInfo)  {
	sendOrder := new(send)
	sendOrder.send(out, params)
}


// receive
type recIter interface {
	rec(in <-chan request.OrderInfo)
}

type rec struct {

}

func NewRecOrder() *rec {
	return &rec{}
}

func (s *rec) rec(in <-chan request.OrderInfo) {
	for {
		select {
		case item := <-in:
			// save to db or other operations...
			fmt.Printf("Id=%d, Sn=%s \n", item.Id, item.Sn)
		}
	}

	//for item := range in {
	//	// save to db or other operations...
	//	fmt.Printf("Id=%d, Sn=%s \n", item.Id, item.Sn)
	//}
}

func (s *rec) HandleOrder(in <-chan request.OrderInfo)  {
	recOrder := new(rec)
	recOrder.rec(in)
}
