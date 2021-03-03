package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostMethod(c *gin.Context) {
	fmt.Println("\napi.go 'PostMethod' called")
	message := "PostMethod called"
	c.JSON(http.StatusOK, message)
}

func GetMethod(c *gin.Context) {
	fmt.Println("\napi.go 'GetMethod' called")
	message := "GetMethod called"
	c.JSON(http.StatusOK, message)
}

func main() {
	router := gin.Default()

	router.POST("/", PostMethod)
	router.GET("/", GetMethod)

	listenPort := "6543"

	router.Run(":" + listenPort)
}
