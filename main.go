package main

import (
	"fmt"
	lesson1 "lesson_1"
)

func main() {
	list := lesson1.New(1)
	list.Add(lesson1.NewNode(10))
	list.Print()
	s := []int{1, 2, 3, 4, 5}
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
