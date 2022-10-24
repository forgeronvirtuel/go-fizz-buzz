package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := setupRouter()
	if err := router.Run("localhost:9080"); err != nil {
		log.Fatalln(err)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.PUT("/int/first", putIntFirst)
	return router
}

type intValue struct {
	Value int `json:"value" binding:"required,min=1"`
}

func putIntFirst(c *gin.Context) {
	log.Println("OK")
	body := intValue{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	log.Printf("Value = %d", body.Value)
	c.JSON(http.StatusOK, &body)
}
