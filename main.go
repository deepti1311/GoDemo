package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	//fmt.Println("hello Go...")
	r := gim.Defalt()

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello REST API")
	})
}
