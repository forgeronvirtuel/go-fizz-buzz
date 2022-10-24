package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
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
	router.GET("/result", fizzbuzz(&int1, &int2, &limit, &s1, &s2))
	return router
}

func format(sb *strings.Builder, i1, i2, i int, s1, s2, concat string) {
	if i%i1 == 0 && i%i2 == 0 {
		sb.WriteString(fmt.Sprintf("%s", concat))
	} else if i%i1 == 0 {
		sb.WriteString(fmt.Sprintf("%s", s1))
	} else if i%i2 == 0 {
		sb.WriteString(fmt.Sprintf("%s", s2))
	} else {
		sb.WriteString(fmt.Sprintf("%d", i))
	}
}

func fizzbuzz(i1, i2, limit *int, s1, s2 *string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var sb strings.Builder
		concat := fmt.Sprintf("%s%s", *s1, *s2)
		for i := 1; i < *limit; i++ {
			format(&sb, *i1, *i2, i, *s1, *s2, concat)
			sb.WriteRune(',')
		}
		format(&sb, *i1, *i2, *limit, *s1, *s2, concat)
		c.String(http.StatusOK, sb.String())
	}
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
		c.Status(http.StatusOK)
	}
}

func createRouteGetInt(varname string, value *int) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, formatIntvalue, varname, *value)
	}
}
