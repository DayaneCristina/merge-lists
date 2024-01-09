package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	L "americanas.teste/lists"
)

func main() {
	r := gin.Default()

	r.POST("/saveLists", L.SaveLists)

	r.GET("/merge", L.Merge)

	fmt.Println("Server running on port 8080")

	r.Run(":8080")
}
