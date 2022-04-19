package main

import "github.com/gin-gonic/gin"
import 	"github.com/abhinav054m/keycloak-gin-auth/keycloak"

func main() {
	r := gin.Default()
	r.Use(keycloak.Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

