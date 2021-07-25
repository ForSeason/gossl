//
// lru.go
// Copyright (C) 2021 forseason <me@forseason.vip>
//
// Distributed under terms of the MIT license.
//

// This package provides implement of lru using
// `doublylinkedlist` package and golang's map
package lru

import (
	"sync"

	"github.com/forseason/gossl/doublylinkedlist"
	"github.com/forseason/gossl/doublylinkedlist/node"
)

type LRU struct {
	m              map[interface{}]*node.Node
	l              *doublylinkedlist.DoublyLinkedList
	size, capacity int
	mutex          sync.Mutex
}

type KVPair struct {
	k, v interface{}
}

// New(capacity int) creates a LRU cache with given capacity.
// Parameter `capacity` limits the max amount of data. If LRU
// is full, new value will be instead of oldest value.
func New(capacity int) *LRU {
	return &LRU{
		m:        make(map[interface{}]*node.Node),
		l:        doublylinkedlist.New(),
		capacity: capacity,
		size:     0,
	}
}

func (this *LRU) Store(key, value interface{}) bool {
	if key == nil || value == nil {
		return false
	}

	this.mutex.Lock()
	defer this.mutex.Unlock()

	if _, ok := this.m[key]; ok {
		return false
	}

	data := &KVPair{
		k: key,
		v: value,
	}
	if this.capacity > this.size {
		node := this.l.AddBeforeTail(data)
		this.m[key] = node
		this.size++
	} else {
		oldNode := this.l.GetHead()
		delete(this.m, oldNode.Val.(*KVPair).k)
		this.l.RemoveAfterHead()
		node := this.l.AddBeforeTail(data)
		this.m[key] = node
	}

	return true
}

func (this *LRU) Get(key interface{}) (interface{}, bool) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if _, ok := this.m[key]; !ok {
		return nil, false
	} else {
		node := this.m[key]
		this.l.RemoveNode(node)
		this.l.AddBeforeTail(node)
		return node.Val.(*KVPair).v, true
	}
}

func (this *LRU) Delete(key interface{}) bool {
	if _, ok := this.m[key]; !ok {
		return false
	}
	this.l.RemoveNode(this.m[key])
	delete(this.m, key)
	return true
}

func (this *LRU) GetSize() int {
	return this.size
}

func (this *LRU) GetCapacity() int {
	return this.capacity
}

// Reserve(capacity int) will reset the capacity of LRU if
// the new value is bigger than the old one.
func (this *LRU) Reserve(capacity int) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if this.capacity >= capacity {
		return false
	}
	this.capacity = capacity
	return true
}
