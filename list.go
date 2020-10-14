package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(13)
	intList.PushBack(30)

	fmt.Printf("\nLIST: %v\n", intList)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Printf("Element: %d\n", element.Value.(int))
	}
}
