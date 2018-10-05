package main

import (
	"fmt"
)

type(
 	linkList struct {
 		left *nod
		right  *nod
		lenth int
		nil *nod
	}
 	nod struct {
 		k interface{}
 		prev *nod
 		next *nod
	}
)
func  NewLinkList() *linkList  {
	return &linkList{
		nil,
		nil,
		0,
		nil,
	}
}
//左添加
func (link *linkList)Lappend(v interface{})  {
	if link.lenth==0 {
		n:= &nod{v,link.nil,link.nil}
		//link.tail=n
		link.left=n
		link.right=n
	}else {
		n:= &nod{v,link.nil,link.left}
		link.left.prev=n
		link.left=n
	}
	link.lenth++
	fmt.Print(link.lenth)
}
//右添加
func (link *linkList)Rappend(v interface{})  {
	if link.lenth==0 {
		n:= &nod{v,link.nil,link.nil}
		link.left=n
		link.right=n
	}else {
		n:= &nod{v,link.right,nil}
		link.right.next=n
		link.right=n
	}
	link.lenth++
	fmt.Print(link.lenth)

}
//查询
func (link *linkList)queryl(v interface{}) *nod{
	if link.lenth == 0 {
		return nil
	} else  {

		for node := link.left; node!=nil; node=node.next{
			if node.k == v {
				return node
			}
		}
	}
	return nil
}
//删除
func (link *linkList)deletee(v interface{}) interface{} {
	node:=link.queryl(v)
	if node!=nil {
		if node.prev == nil {
			link.left = link.left.next
		} else {
			node.prev.next=node.next
		}
		if node.next == nil {
			link.right = link.right.prev
		}else {
			node.next.prev= node.prev
		}
		link.lenth--
	}
	return 0
}
func main() {

	link:=NewLinkList()
	link.Lappend("你")
	//link.Lappend("我")
	//link.Lappend("他")
	link.Rappend("谁")

	fmt.Printf("%+v", link.queryl("你"))
	fmt.Printf("%+v", link.queryl("谁"))

	link.deletee("你")
	link.deletee("谁")

	fmt.Print(link.lenth)
}
