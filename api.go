package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var savedList1 *ListNode
var savedList2 *ListNode

func saveLists(c *gin.Context) {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{5, nil}}}
	savedList1 = list1
	savedList2 = list2

	c.JSON(200, gin.H{
		"message": "Lists saved successfully",
	})
}

func mergeLists(c *gin.Context) {
	if savedList1 == nil || savedList2 == nil {
		c.JSON(400, gin.H{
			"error": "Lists not saved yet",
		})
		return
	}

	result := mergeTwoLists(savedList1, savedList2)

	mergedValues := []int{}
	for result != nil {
		mergedValues = append(mergedValues, result.Val)
		result = result.Next
	}

	c.JSON(200, gin.H{
		"mergedList": mergedValues,
	})
}

func main() {
	r := gin.Default()

	r.POST("/saveLists", saveLists)
	r.GET("/merge", mergeLists)

	fmt.Println("Server running on port 8080")
	r.Run(":8080")
}
