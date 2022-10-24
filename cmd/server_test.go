package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetFirstValue(t *testing.T) {
	router := setupRouter()

	// Read the first int, default value
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/?int1=3&int2=5&limit=15&str1=fizz&str2=buzz", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz", w.Body.String())
}
