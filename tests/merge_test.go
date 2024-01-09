package main

import (
	"reflect"
	"testing"
	L "americanas.teste/lists"
	S "americanas.teste/structs"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := L.ArrayToList([]int{1, 2, 4})
	list2 := L.ArrayToList([]int{1, 3, 5})

	expected := []int{1, 1, 2, 3, 4, 5}

	result := L.MergeTwoLists(list1, list2)

	if !validateResult(result, expected) {
		t.Errorf("Resultado incorreto. Esperado: %v, Obtido: %v", expected, getListValues(result))
	}

	list3 := L.ArrayToList([]int{2, 3, 4})
	list4 := L.ArrayToList([]int{1, 5, 6})

	expected = []int{1, 2, 3, 4, 5, 6}

	result = L.MergeTwoLists(list3, list4)

	if !validateResult(result, expected) {
		t.Errorf("Resultado incorreto. Esperado: %v, Obtido: %v", expected, getListValues(result))
	}
}

func validateResult(result *S.ListNode, expected []int) bool {
	values := getListValues(result)
	return reflect.DeepEqual(values, expected)
}

func getListValues(node *S.ListNode) []int {
	values := []int{}
	for node != nil {
		values = append(values, node.Val)
		node = node.Next
	}
	return values
}
