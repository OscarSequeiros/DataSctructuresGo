package main

import (
	"DataStructuresGolang/linkedlist"
	"DataStructuresGolang/stack"
	"fmt"
)

func main() {

	emptyList := linkedlist.BuildEmptyLinkedList()
	list := emptyList.Push(1).Append(6).Push(2).Push(3).Append(4).Append(5)
	nodeAt3 := list.NodeAt(0)

	fmt.Printf("%v with size: %v", list.String(), list.Size)
	//list.Insert(8, nodeAt3)
	println()
	fmt.Printf("element removed %v and list: %v", list.RemoveAfter(nodeAt3), list.String())
	fmt.Printf("reversed list : %v", list.Reverse())
	println()

	aStack := stack.CreateFrom([]string{"a", "b", "c", "d"})
	println(aStack.String())

	expression := "(2 + 2) - (1 + 2) * 3"
	println(stack.CheckParentheses(expression))
}
