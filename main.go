package main

import (
	"lesson_1"
	"fmt"
)


func main() {
	list := lesson1.New(1)
	list.Add(&lesson1.Node{Val: 2, Next: nil})
	list.Print()
	s := [] int {1, 2, 3, 4, 5}
	list = lesson1.NewFromSlice(s)
	list.Print()
	fmt.Println(list.At(3))
	list.Pop()
	list.Print()
	list.UpdateAt(3, 1)
	list.Print()
	list.DeleteFrom(2)
	list.Print()
	fmt.Println(list.Size())
}