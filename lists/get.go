package lists

import (
	"errors"
)

func GetMerged() ([]int, error) {
	if savedList1 == nil || savedList2 == nil {
		return []int{}, errors.New("Listas não encontradas")
	}

	list1 := ArrayToList(savedList1)
	list2 := ArrayToList(savedList2)

	result := MergeTwoLists(list1, list2)

	mergedValues := []int{}

	for result != nil {
		mergedValues = append(mergedValues, result.Val)
		result = result.Next
	}

	return mergedValues, nil
}
