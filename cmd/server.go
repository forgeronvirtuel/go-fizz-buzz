package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	int1name    = "int1"
	int2name    = "int1"
	formatvalue = "%s = %d"
)

func main() {
	router := setupRouter()
	if err := router.Run("localhost:9080"); err != nil {
		log.Fatalln(err)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	int1 := 0
	int2 := 0
	router.GET("/int/first", createRouteGetInt(int1name, &int1))
	router.PUT("/int/first", createRoutePutInt(&int1))
	router.GET("/int/second", createRouteGetInt(int2name, &int2))
	router.PUT("/int/second", createRoutePutInt(&int2))
	return router
}

type intValue struct {
	Value int `json:"value" binding:"required,min=1"`
}

func createRoutePutInt(value *int) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := intValue{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		*value = body.Value
		c.JSON(http.StatusOK, &body)
	}
}

func createRouteGetInt(varname string, value *int) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, formatvalue, varname, *value)
	}
}
