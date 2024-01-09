package lists

import (
	"github.com/gin-gonic/gin"
    "net/http"
	S "americanas.teste/structs"
)

func SaveLists(c *gin.Context) {

	var bodyRequest S.RequestSave

	if err := c.BindJSON(&bodyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao processar o JSON",
		})
		return
	}

	Save(bodyRequest.List1, bodyRequest.List2)

	c.JSON(http.StatusOK, gin.H{
		"message": "Listas salvas com sucesso",
	})
}

func Merge(c *gin.Context) {
	result, err := GetMerged()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Listas não encontradas",
		})
		return
	}

	c.JSON(200, gin.H{
		"mergedList": result,
	})
}