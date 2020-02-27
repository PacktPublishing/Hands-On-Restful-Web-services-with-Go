package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	/* GET takes a route and a handler function
	   Handler takes the gin context object
	*/
	r.GET("/pingTime", func(c *gin.Context) {
		// JSON serializer is available on gin context
		c.JSON(200, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})

	r.Run(":8000") // Listen and serve on 0.0.0.0:8080
}
