package main

import (
	"fmt"
)

type(
	queue struct {
		head *nodee
		tail *nodee
		lenth int
	}
	nodee struct {
		value interface{}
		next *nodee

	}
)

func NewQueue() *queue  {
	return &queue{
		nil,
		nil,
		0,
	}
}
func (q *queue)Inqueue(value interface{})  {
	n := &nodee{value,nil}

	if q.lenth==0 {
		q.head = n
		q.tail=q.head

	}else {
		q.tail.next=n
		q.tail = n
	}
	q.lenth++


}
func (q *queue)Dequeue() interface{}  {
	if q.lenth==0 {
		return nil
	}
	deV:=q.head.value
	if q.lenth==1{
		q.head=nil
		q.tail=nil
	}else {
		q.head=q.head.next
	}
	q.lenth--
	return deV
}
func main() {
	queue:=NewQueue()
	queue.Inqueue(1)
	queue.Inqueue(2)
	queue.Inqueue(3)

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

}