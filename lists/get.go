package lists

import (
	"errors"
)

func GetMerged() ([]int, error) {
	if savedList1 == nil || savedList2 == nil {
		return []int{}, errors.New("Listas não encontradas")
	}

	result := MergeTwoLists(savedList1, savedList2)

	mergedValues := []int{}

	for result != nil {
		mergedValues = append(mergedValues, result.Val)
		result = result.Next
	}
	
	return mergedValues, nil
}
