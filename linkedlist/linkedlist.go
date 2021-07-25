//
// linkedlist.go
// Copyright (C) 2021 forseason <me@forseason.vip>
//
// Distributed under terms of the MIT license.
//

// This package provides an implement of linkedlist.
package linkedlist

import (
	"github.com/forseason/gossl/linkedlist/node"
)

// Defination of struct Linkedlist.
type LinkedList struct {
	head *node.Node
}

func New() *LinkedList {
	// Attention: the head node should never be used.
	return &LinkedList{
		head: node.New(nil),
	}
}

func (this *LinkedList) AddAtHead(value interface{}) {
	newNode := node.New(value)
	newNode.Next = this.head.Next
	this.head.Next = newNode
}

func (this *LinkedList) AddAtTail(value interface{}) {
	newNode := node.New(value)
	prevNode := this.head
	for prevNode.Next != nil {
		prevNode = prevNode.Next
	}
	prevNode.Next = newNode
}

// AddAtIndex(index int, value interface{}) (bool) will insert a new
// node behind the node with given index. If index is negative, a new
// node will be inserted in the beginning.
func (this *LinkedList) AddAtIndex(index int, value interface{}) bool {
	if index <= 0 {
		this.AddAtHead(value)
		return true
	}
	prevNode := this.head
	for index > 0 {
		if prevNode.Next == nil {
			return false
		}
		prevNode, index = prevNode.Next, index-1
	}
	newNode := node.New(value)
	newNode.Next = prevNode.Next
	prevNode.Next = newNode
	return true
}

func (this *LinkedList) DeleteAtIndex(index int) bool {
	if index < 0 {
		return false
	}
	prevNode := this.head
	for index > 0 {
		if prevNode.Next == nil {
			return false
		}
		prevNode, index = prevNode.Next, index-1
	}
	prevNode.Next = prevNode.Next.Next
	return true
}

func (this *LinkedList) Get(index int) (interface{}, bool) {
	if index < 0 {
		return nil, false
	}
	currNode := this.head.Next
	for index > 0 {
		if currNode == nil {
			return nil, false
		}
		currNode, index = currNode.Next, index-1
	}
	return currNode.Val, true
}
