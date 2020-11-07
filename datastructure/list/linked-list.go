package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}
type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(property int) {
	newNode := Node{property: property}
	if linkedList.headNode != nil {
		newNode.nextNode = linkedList.headNode
	}
	linkedList.headNode = &newNode
}

func (linkedList *LinkedList) AddToEnd(property int) {
	lastNode := linkedList.LastNode()
	lastNode.nextNode = &Node{property: property}
}

func (linkedList *LinkedList) LastNode() *Node {
	lastNode := &Node{}
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (linkedList *LinkedList) IterateList() {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println("Node's property: ", node.property)
	}
}

func main() {
	linkedList := LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToEnd(10)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.IterateList()
}
