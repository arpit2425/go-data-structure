package linkedlist

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}
type LinkedList interface {
	AddToStart(Val int)
	GetHead() int
	AddToEnd(Val int)
	Iterate()
	GetTail() *Node
	IsEmpty() bool
	RemoveFromEnd()
	RemoveFromFront()
}
type linkedList struct {
	head *Node
}

func NewLinkedList() LinkedList {
	return &linkedList{}
}

func (li *linkedList) AddToStart(Val int) {
	if li.head == nil {
		li.head = &Node{Val, nil}
		return
	} else {
		node := &Node{Val, nil}
		temp := li.head
		node.Next = temp
		li.head = node
	}
}
func (li *linkedList) GetHead() int {
	if li.head != nil {
		return li.head.Val
	}
	return 0
}

func (li *linkedList) IsEmpty() bool {
	return li.head == nil
}
func (li *linkedList) GetTail() *Node {
	if li.IsEmpty() {
		return nil
	}
	temp := li.head
	for ; temp.Next != nil; temp = temp.Next {
	}
	return temp

}

func (li *linkedList) Iterate() {
	temp := li.head
	for ; temp != nil; temp = temp.Next {
		fmt.Print(temp.Val, " -> ")
	}
	fmt.Println()
}

func (li *linkedList) AddToEnd(Val int) {
	if li.head == nil {
		li.head = &Node{Val, nil}
		return
	}
	tail := li.GetTail()
	tail.Next = &Node{Val, nil}
}

func (li *linkedList) RemoveFromEnd() {
	if !li.IsEmpty() {
		fast := li.head.Next
		slow := li.head
		if fast == nil {
			li.head = nil
			return
		}
		for ; fast.Next != nil; slow = slow.Next {
			fast = fast.Next
		}
		slow.Next = nil
	}
}

func (li *linkedList) RemoveFromFront() {
	if !li.IsEmpty() {
		if li.head.Next == nil {
			li.head = nil
			return
		}
		li.head = li.head.Next
	}
}
