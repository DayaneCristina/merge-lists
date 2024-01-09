package lists

import (
	S "americanas.teste/structs"
)

func MergeTwoLists(list1 []int, list2 []int) *S.ListNode {
	head1 := ArrayToList(list1)
	head2 := ArrayToList(list2)

	dummy := &S.ListNode{}
	current := dummy

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			current.Next = head1
			head1 = head1.Next
		} else {
			current.Next = head2
			head2 = head2.Next
		}

		current = current.Next
	}

	if head1 != nil {
		current.Next = head1
	} else {
		current.Next = head2
	}

	return dummy.Next
}
