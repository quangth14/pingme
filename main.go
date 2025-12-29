package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Hello, World!")

	r := gin.Default()
	r.POST("/ping", func(c *gin.Context) {
		// print header
		headersJSON, _ := json.MarshalIndent(c.Request.Header, "", "  ")
		log.Printf("Request Headers:\n%s\n", string(headersJSON))

		// print request body json pretty
		var body map[string]interface{}
		if err := c.ShouldBindJSON(&body); err != nil {
			log.Printf("Request Body Error: %s\n", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		bodyJSON, _ := json.MarshalIndent(body, "", "  ")
		log.Printf("Request Body:\n%s\n", string(bodyJSON))

		c.String(http.StatusOK, "pong")
	})

	r.GET("/.well-known/apple-app-site-association", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.File("./static/.well-known/apple-app-site-association")
	})

	r.Run(":8080")
}
