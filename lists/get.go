package lists

import (
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	if savedList1 == nil || savedList2 == nil {
		c.JSON(400, gin.H{
			"error": "Lists not saved yet",
		})
		return
	}

	result := MergeTwoLists(savedList1, savedList2)

	mergedValues := []int{}

	for result != nil {
		mergedValues = append(mergedValues, result.Val)
		result = result.Next
	}

	c.JSON(200, gin.H{
		"mergedList": mergedValues,
	})
}
