//
// doublylinkedlist.go
// Copyright (C) 2021 forseason <me@forseason.vip>
//
// Distributed under terms of the MIT license.
//

//This package provides an implement of doubly-linkedlist
package doublylinkedlist

import (
	"github.com/forseason/gossl/doublylinkedlist/node"
)

type DoublyLinkedList struct {
	head *node.Node
	tail *node.Node
}

func New() *DoublyLinkedList {
	//Attention: head and tail should never be used!!
	list := &DoublyLinkedList{
		head: node.New(nil),
		tail: node.New(nil),
	}
	list.head.Next = list.tail
	list.tail.Prev = list.head
	return list
}

func (this *DoublyLinkedList) AddAfterHead(value interface{}) *node.Node {
	p := node.New(value)
	p.Next = this.head.Next
	p.Prev = this.head
	this.head.Next.Prev = p
	this.head.Next = p
	return p
}

func (this *DoublyLinkedList) AddBeforeTail(value interface{}) *node.Node {
	p := node.New(value)
	p.Prev = this.tail.Prev
	p.Next = this.tail
	this.tail.Prev.Next = p
	this.tail.Prev = p
	return p
}

func (this *DoublyLinkedList) RemoveAfterHead() bool {
	if this.head.Next == this.tail {
		return false
	}

	this.head.Next = this.head.Next.Next
	this.head.Next.Next.Prev = this.head
	return true
}

func (this *DoublyLinkedList) RemoveBeforeTail() bool {
	if this.head.Next == this.tail {
		return false
	}

	this.tail.Prev = this.tail.Prev.Prev
	this.tail.Prev.Prev.Next = this.tail
	return true
}

func (this *DoublyLinkedList) RemoveNode(node *node.Node) bool {
	if node == this.head || node == this.tail {
		return false
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	return true
}

func (this *DoublyLinkedList) GetHead() *node.Node {
	if this.head.Next == this.tail {
		return nil
	}
	return this.head.Next
}

func (this *DoublyLinkedList) GetTail() *node.Node {
	if this.head.Next == this.tail {
		return nil
	}
	return this.tail.Prev
}

func (this *DoublyLinkedList) Get(index int) *node.Node {
	if index < 0 {
		return nil
	}
	p := this.head.Next
	for p != this.tail {
		if index == 0 {
			return p
		}
		p = p.Next
		index--
	}
	return nil
}
