package main

import (
	"fmt"
	L "americanas.teste/lists"
)

func main() {
	list1 := []int{1,2,4}
	list2 := []int{1,3,5}

	result := L.MergeTwoLists(list1, list2)

	fmt.Println("Lista mesclada:")
	
	for result != nil {
		fmt.Printf("%d", result.Val)

		if (result.Next != nil ) {
			fmt.Printf(" -> ")
		}

		result = result.Next
	}
}
