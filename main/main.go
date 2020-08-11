package main

import (
	"fmt"
	"github.com/hubertme/gin-demo/gin"
)

func main() {
	fmt.Println("Hello, world!")

	gin.SetupGinServer()
}