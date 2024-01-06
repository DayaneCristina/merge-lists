package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // Nó fictício para começar a lista mesclada
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// Se uma das listas terminou, anexa o restante da outra lista
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next // Retorna a lista mesclada, ignorando o nó fictício inicial
}

func main() {
	// Exemplo de uso
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{5, nil}}}

	result := mergeTwoLists(list1, list2)

	fmt.Println("Lista mesclada:")
	for result != nil {
		fmt.Printf("%d -> ", result.Val)
		result = result.Next
	}
	fmt.Println("nil")
}
