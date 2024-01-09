package lists

import (
	"github.com/gin-gonic/gin"
	S "americanas.teste/structs"
)

var savedList1 []int
var savedList2 []int

func SaveLists(c *gin.Context) {
	var bodyRequest S.RequestSave

	if err := c.BindJSON(&bodyRequest); err != nil {
		c.JSON(400, gin.H{
			"message": "Erro ao processar o JSON",
		})
		return
	}

	savedList1 = bodyRequest.List1
	savedList2 = bodyRequest.List2

	c.JSON(200, gin.H{
		"message": "Lists saved successfully",
	})
}