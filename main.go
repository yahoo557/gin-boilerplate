package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yahoo557/gin-boilerplate/internal/user"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		c.Writer.Write([]byte("latency: " + latency.String()))
	}
}

func main() {
	r := gin.Default()
	r.Use(Logger())

	apiV1 := r.Group("/api/v1")

	user.UserRouter(apiV1)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
