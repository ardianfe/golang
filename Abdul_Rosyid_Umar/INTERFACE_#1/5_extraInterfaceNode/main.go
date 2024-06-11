package main

import "fmt"

// mendefinisikan Node interface
type Node interface {
	Next() *ListNode
}

// Struct yang merepresentasikan Node
type ListNode struct {
	Data     int
	NextNode *ListNode
}

// Method untuk return node selanjutnya
func (l ListNode) Next() *ListNode {
	return l.NextNode
}

func main() {
	// membuat list nodes
	node1 := ListNode{Data: 10}
	node2 := ListNode{Data: 20}
	node3 := ListNode{Data: 30}
	node1.NextNode = &node3
	node2.NextNode = &node3

	// Assign node ke Node interface
	var n Node
	n = node1

	// memanggil Next method via Node interface
	nextNode := n.Next()
	fmt.Println("Next Node Data:", nextNode.Data)
}
