package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func main() {
	// logrus.SetLevel(logrus.DebugLevel)

	// pprof.StartPprof()

	rand.Seed(time.Now().UnixNano())

	bindAddress := "0.0.0.0:2303"
	r := gin.Default()
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(bindAddress)
}
