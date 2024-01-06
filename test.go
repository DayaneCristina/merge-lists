package main

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	// 1 caso de teste
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{5, nil}}}

	expected := []int{1, 1, 2, 3, 4, 5}

	result := mergeTwoLists(list1, list2)

	// Verifica a lista mesclada se está correta
	if !validateResult(result, expected) {
		t.Errorf("Resultado incorreto. Esperado: %v, Obtido: %v", expected, getListValues(result))
	}

	// 2 caso teste
	list3 := &ListNode{2, &ListNode{3, &ListNode{4, nil}}}
	list4 := &ListNode{1, &ListNode{5, &ListNode{6, nil}}}

	expected = []int{1, 2, 3, 4, 5, 6}

	result = mergeTwoLists(list3, list4)

	// Verifica a lista mesclada se está correta
	if !validateResult(result, expected) {
		t.Errorf("Resultado incorreto. Esperado: %v, Obtido: %v", expected, getListValues(result))
	}
}

// Função auxiliar para validar se a lista mesclada está correta
func validateResult(result *ListNode, expected []int) bool {
	values := getListValues(result)
	return reflect.DeepEqual(values, expected)
}

// Função auxiliar para obter os valores da lista mesclada
func getListValues(node *ListNode) []int {
	values := []int{}
	for node != nil {
		values = append(values, node.Val)
		node = node.Next
	}
	return values
}
