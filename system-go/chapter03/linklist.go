package chapter03

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func add(head *Node, val int) *Node {
	if head == nil {
		head = &Node{val, nil}
	} else {
		var temp *Node
		temp = head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &Node{val, nil}
	}
	return head
}

func traverse(head *Node) {
	for head != nil {
		fmt.Println(head.Value)
		head = head.Next
	}
}

func Start04() {
	// adding elements in link list
	var head = new(Node)
	head = nil
	for i := 0; i < 10; i++ {
		head = add(head, i)
	}

	// traverse the list
	traverse(head)
}
