package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetFirstValue(t *testing.T) {
	router := setupRouter()

	// Read the first int, default value
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/int/first", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(formatvalue, int1name, 0), w.Body.String())

	// Read the second int, default value
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/int/second", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(formatvalue, int2name, 0), w.Body.String())

	// Set the first value to 3
	body := intValue{Value: 3}
	v, _ := json.Marshal(body)
	req, _ = http.NewRequest(http.MethodPut, "/int/first", bytes.NewReader(v))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Set the second value to 5
	body = intValue{Value: 5}
	v, _ = json.Marshal(body)
	req, _ = http.NewRequest(http.MethodPut, "/int/second", bytes.NewReader(v))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Read the first int, default value
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/int/first", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(formatvalue, int1name, 3), w.Body.String())

	// Read the second int, default value
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/int/second", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(formatvalue, int2name, 5), w.Body.String())

}
