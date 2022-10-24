package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type fizzBuzzParam struct {
	Int1  int    `form:"int1" binding:"required,min=1"`
	Int2  int    `form:"int2" binding:"required,min=1"`
	Limit int    `form:"limit" binding:"required,min=1"`
	Str1  string `form:"str1" binding:"required"`
	Str2  string `form:"str2" binding:"required"`
}

func FizzbuzzRoute(c *gin.Context) {
	var params fizzBuzzParam
	if err := c.Bind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, fizzbuzz(params.Int1, params.Int2, params.Limit, params.Str1, params.Str2))
}

func fizzbuzz(i1, i2, limit int, s1, s2 string) string {
	var sb strings.Builder
	concat := fmt.Sprintf("%s%s", s1, s2)
	for i := 1; i < limit; i++ {
		format(&sb, i1, i2, i, s1, s2, concat)
		sb.WriteRune(',')
	}
	format(&sb, i1, i2, limit, s1, s2, concat)
	return sb.String()
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
