package main

import (
	"fmt"
	"github.com/hubertme/gin-demo/database"
	"github.com/hubertme/gin-demo/gin"
)

func main() {
	fmt.Println("Hello, world!")
	database.InitDriver()

	err := gin.SetupGinServer()

	if err != nil {
		fmt.Println("Error in running gin server:", err.Error())
	}
}
