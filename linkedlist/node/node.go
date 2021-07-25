//
// node.go
// Copyright (C) 2021 forseason <me@forseason.vip>
//
// Distributed under terms of the MIT license.
//

// This package provides implement of node.
package node

// defination of node for linkedlist
type Node struct {
	Val  interface{}
	Next *Node
}

// New(value interface{}) *Node accepts an interface as value
// and returns a pointer. The Next field will be set as nil.
func New(value interface{}) *Node {
	return &Node{
		Val:  value,
		Next: nil,
	}
}
