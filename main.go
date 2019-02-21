/*
lru算法
 */

package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type LruTable struct {
	capcity int
	table []*Link
	head *LinkNode
	lruCapcity int
	lruLimit int
}
type Link struct {
	head *LinkNode
	tail *LinkNode
}

type LinkNode struct {
	data int
	next *LinkNode
	lruNext *LinkNode
}

func main() {
	const count=2
	table:=LruTable{}
	table.init(count)
	data:=make([]int,0)
	for i:=0;i<count;i++{
		a:=rand.Intn(100)
		table.add(&LinkNode{data:a})
		data=append(data,a)
	}
	fmt.Println(data)
	table.show()
}
func (link *Link)init(){
	link.head=&LinkNode{}
	link.tail=link.head
}
func (table *LruTable)init(capcity int){
	table.capcity=capcity
	table.table=make([]*Link,capcity)
	table.lruCapcity=capcity/2
	if table.lruCapcity<2{
		panic(errors.New("lru 长度必须大于1"))
	}
}
func (table *LruTable)add(node *LinkNode){
	idx:=table.hash(node.data)
	head:=table.table[idx]
	if head==nil{
		link:=&Link{}
		link.init()
		table.table[idx]=link
		link.tail.next=node
		link.tail=link.tail.next
		table.lruLimit=table.lruLimit+1
	}else{
		//查找该节点返是否已经存在
		headNode:=head.head
		for headNode.next!=nil&&headNode.next.data!=node.data{
			headNode=headNode.next
		}
		//如果未找到给定的节点，将该节点添加到该链表的末尾
		if headNode.next==nil{
			head.tail.next=node
			head.tail=head.tail.next
			table.lruLimit=table.lruLimit+1
		}
	}
	//将lruhead指向最新添加的节点
	node.lruNext=table.head
	table.head=node
	//删除lru多余的节点
	if table.lruLimit>table.lruCapcity{
		head:=table.head
		for head.lruNext.lruNext!=nil{
			head=head.lruNext
		}
		head.lruNext=nil
		table.lruLimit=table.lruLimit-1
	}
}
func (table *LruTable)show(){
	head:=table.head
	for head!=nil{
		fmt.Print(head.data," ")
		head=head.lruNext
	}
	fmt.Println()
}
func (table *LruTable)hash(data int)int{
	return data%table.capcity
}