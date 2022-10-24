package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	limitname         = "limit"
	int1name          = "int1"
	int2name          = "int1"
	string1name       = "str1"
	string2name       = "str2"
	formatIntvalue    = "%s = %d"
	formatStringvalue = "%s = %s"
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
	limit := 0
	s1 := ""
	s2 := ""
	router.GET("/limit", createRouteGetInt(limitname, &limit))
	router.PUT("/limit", createRoutePutInt(&limit))
	router.GET("/int/first", createRouteGetInt(int1name, &int1))
	router.PUT("/int/first", createRoutePutInt(&int1))
	router.GET("/int/second", createRouteGetInt(int2name, &int2))
	router.PUT("/int/second", createRoutePutInt(&int2))
	router.GET("/string/first", createRouteGetString(string1name, &s1))
	router.PUT("/string/first", createRoutePutString(&s1))
	router.GET("/string/second", createRouteGetString(string2name, &s2))
	router.PUT("/string/second", createRoutePutString(&s2))
	return router
}

type stringValue struct {
	Value string `json:"value" binding:"required"`
}

func createRoutePutString(value *string) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := stringValue{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		*value = body.Value
		c.JSON(http.StatusOK, &body)
	}
}

func createRouteGetString(varname string, value *string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, formatStringvalue, varname, *value)
	}
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
		c.String(http.StatusOK, formatIntvalue, varname, *value)
	}
}
