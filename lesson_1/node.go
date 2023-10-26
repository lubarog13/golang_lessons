package lesson1
import ("fmt")

type Node struct {
	Val   int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func New(q int) *LinkedList {
	list := &LinkedList{Head: nil}
	for i := 0; i < q; i++ {
		list.Add(&Node{Val: 0, Next: nil})
	}
	return list
}


func (l *LinkedList) Add(n *Node) {
	node := l.Head
	if node == nil { 
		l.Head = n
		return 
	}
	for node.Next != nil {
		node = node.Next
	}
	node.Next = n
}

func (l *LinkedList) Pop() {
	node := l.Head
	if node == nil {
		return
	}
	if node.Next == nil {
		l.Head = nil
		return
	}
	last_node := node
	for node.Next != nil {
		last_node = node
		node = node.Next
	}
	last_node.Next = nil
}

func (l *LinkedList) At(pos int) int {
	node := l.Head
	counter := 0
	for node != nil {
		if (counter == pos) {
			return node.Val
		}
		node = node.Next
		counter++
	}
	return nil
}

func (l *LinkedList) Size() int {
	node := l.Head
	counter := 0
	for node != nil {
		counter++
		node = node.Next
	}
	return counter
}

func (l *LinkedList) DeleteFrom(pos int) {
	node := l.Head
	if node == nil {
		return
	}
	if pos == 0 && node.Next != nil {
		l.Head = node.Next
		return
	}
	if pos == 0 {
		l.Head = nil
		return
	}
	last_node := node
	counter := 0
	for node != nil {
		last_node = node
		node = node.Next
		counter++
		if (counter == pos) {
			last_node.Next = node.Next
		}
	}
} 

func (l *LinkedList) UpdateAt(pos, val int) {
	node := l.Head
	counter := 0
	for node != nil {
		if (counter == pos) {
			node.Val = val
			return
		}
		node = node.Next
		counter++
	}
}

func NewFromSlice(s []int) *LinkedList {
	list := &LinkedList{Head: &Node{Val: 0, Next: nil}}
	for _, v := range s {
		list.Add(&Node{Val: v, Next: nil})
	}
	return list
}

func (list *LinkedList) Print() {
	node := list.Head
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	fmt.Println()
}