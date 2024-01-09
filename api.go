package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	L "americanas.teste/lists"
)

func main() {
	r := gin.Default()

	r.POST("/lists", L.SaveLists)
	r.GET("/lists", L.Get)

	fmt.Println("Server running on port 8080")

	r.Run(":8080")
}
