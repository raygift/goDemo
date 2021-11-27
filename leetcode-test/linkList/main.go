package main

import "fmt"

func main() {
	obj := Constructor()
	// ["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
	// [[],[1],[3],[1,2],[1],[1],[1]]

	// 	["MyLinkedList","addAtHead","addAtHead","addAtHead","addAtIndex","deleteAtIndex","addAtHead","addAtTail","get","addAtHead","addAtIndex","addAtHead"]
	// [[],[7],[2],[1],[3,0],[2],[6],[4],[4],[4],[5,0],[6]]

	obj.AddAtHead(7)
	obj.AddAtHead(2)
	obj.AddAtHead(1)
	for p := obj.preHead; p != nil; p = p.Next {
		fmt.Printf("%p %#v\n", p, p)
	}
	fmt.Printf("\n")

	obj.AddAtIndex(3, 0)
	for p := obj.preHead; p != nil; p = p.Next {
		fmt.Printf("%p %#v\n", p, p)
	}
	obj.DeleteAtIndex(2)
	obj.AddAtHead(6)

	obj.AddAtTail(4)

	fmt.Println(obj.Get(4))

	obj.AddAtHead(4)
	obj.AddAtIndex(5, 0)
	obj.AddAtHead(6)
	for p := obj.preHead; p != nil; p = p.Next {
		fmt.Printf("%p %#v\n", p, p)
	}
	fmt.Printf("\n")

}

type MyLinkedList struct {
	size      int
	preHead   *MyListNode
	afterTail *MyListNode
}
type MyListNode struct {
	Val  int
	Next *MyListNode
	Pre  *MyListNode
}

func Constructor() MyLinkedList {
	p := &MyListNode{Val: -1}
	a := &MyListNode{Val: -1}
	p.Next = a
	a.Pre = p
	newList := MyLinkedList{size: 0, preHead: p, afterTail: a}
	return newList
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.preHead
	for j := 0; j <= index; j++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &MyListNode{Val: val} // 创建节点

	newNode.Next = this.preHead.Next
	newNode.Pre = this.preHead
	this.preHead.Next.Pre = newNode
	this.preHead.Next = newNode
	this.size++

}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &MyListNode{Val: val} // 创建节点

	newNode.Next = this.afterTail
	newNode.Pre = this.afterTail.Pre
	this.afterTail.Pre.Next = newNode
	this.afterTail.Pre = newNode
	this.size++

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.size {
		this.AddAtTail(val)
		return
	}
	if index > this.size {
		return
	}
	newNode := &MyListNode{Val: val} // 创建节点

	cur := this.preHead
	for j := 0; j <= index; j++ {
		cur = cur.Next
	}
	newNode.Pre = cur.Pre
	newNode.Next = cur
	cur.Pre.Next = newNode
	cur.Pre = newNode
	this.size++

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}

	cur := this.preHead
	for j := 0; j <= index; j++ {
		cur = cur.Next
	}
	cur.Pre.Next = cur.Next
	cur.Next.Pre = cur.Pre
	cur.Pre = nil
	cur.Next = nil
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
