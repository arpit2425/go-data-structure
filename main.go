package main

import (
	"fmt"
	"go-data-structure/src/linkedlist"
)

func main() {
	RunLinkedList()

}

func RunLinkedList() {
	list := linkedlist.NewLinkedList()
	
	fmt.Println(list.IsEmpty())
	list.AddToStart(10)
	list.RemoveFromFront()
	fmt.Println(list.GetHead())
	list.RemoveFromEnd()

	list.AddToStart(100)
	// fmt.Println(list.GetHead())
	list.AddToStart(40)
	// fmt.Println(list.GetHead())
	list.AddToStart(50)
	list.AddToEnd(3)
	fmt.Println(list.IsEmpty())
	list.Iterate()
	list.RemoveFromEnd()
	tail := list.GetTail()
	fmt.Println(tail.Val)
	list.RemoveFromFront()
	list.RemoveFromFront()
	list.Iterate()
}
