package lists

import(
	"sort"
	S "americanas.teste/structs"
)

func ArrayToList(nums []int) *S.ListNode {
	if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)

	head := &S.ListNode{Val: nums[0]}
	current := head

	for i := 1; i < len(nums); i++ {
		node := &S.ListNode{Val: nums[i]}
		current.Next = node
		current = current.Next
	}

	return head
}