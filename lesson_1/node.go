package lesson1

import (
	"fmt"
)

type node struct {
	val  int
	next *node
}

type linkedList struct {
	head *node
}

func New(q int) *linkedList {
	list := &linkedList{head: nil}
	for i := 0; i < q; i++ {
		list.Add(&node{val: 0, next: nil})
	}
	return list
}

func NewNode(val int) *node {
	return &node{val: val, next: nil}
}

func (l *linkedList) Add(n *node) {
	node := l.head
	if node == nil {
		l.head = n
		return
	}
	for node.next != nil {
		node = node.next
	}
	node.next = n
}

func (l *linkedList) Pop() {
	node := l.head
	if node == nil {
		return
	}
	if node.next == nil {
		l.head = nil
		return
	}
	last_node := node
	for node.next != nil {
		last_node = node
		node = node.next
	}
	last_node.next = nil
}

func (l *linkedList) At(pos int) int {
	node := l.head
	counter := 0
	for node != nil {
		if counter == pos {
			return node.val
		}
		node = node.next
		counter++
	}
	return 0
}

func (l *linkedList) Size() int {
	node := l.head
	counter := 0
	for node != nil {
		counter++
		node = node.next
	}
	return counter
}

func (l *linkedList) DeleteFrom(pos int) {
	node := l.head
	if node == nil {
		return
	}
	if pos == 0 && node.next != nil {
		l.head = node.next
		return
	}
	if pos == 0 {
		l.head = nil
		return
	}
	last_node := node
	counter := 0
	for node != nil {
		last_node = node
		node = node.next
		counter++
		if counter == pos {
			last_node.next = node.next
		}
	}
}

func (l *linkedList) UpdateAt(pos, val int) {
	node := l.head
	counter := 0
	for node != nil {
		if counter == pos {
			node.val = val
			return
		}
		node = node.next
		counter++
	}
}

func NewFromSlice(s []int) *linkedList {
	list := &linkedList{head: &node{val: 0, next: nil}}
	for _, v := range s {
		list.Add(&node{val: v, next: nil})
	}
	return list
}

func (list *linkedList) Print() {
	node := list.head
	for node != nil {
		fmt.Printf("%d ", node.val)
		node = node.next
	}
	fmt.Println()
}
