package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/orders", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{ "message" : "ping" })
  })
  r.Run()
}

