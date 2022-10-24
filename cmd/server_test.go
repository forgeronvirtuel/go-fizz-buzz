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
	assert.Equal(t, fmt.Sprintf(formatIntvalue, int1name, 0), w.Body.String())

	// Read the second int, default value
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/int/second", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(formatIntvalue, int2name, 0), w.Body.String())

	// Read the limit, default value
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/limit", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(formatIntvalue, limitname, 0), w.Body.String())

	// Read the first string, default value
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/string/first", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(formatStringvalue, string1name, ""), w.Body.String())

	// Read the second string, default value
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/string/second", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(formatStringvalue, string2name, ""), w.Body.String())

	// Set the first value to 3
	bodyInt := intValue{Value: 3}
	v, _ := json.Marshal(bodyInt)
	req, _ = http.NewRequest(http.MethodPut, "/int/first", bytes.NewReader(v))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Set the second value to 5
	bodyInt = intValue{Value: 5}
	v, _ = json.Marshal(bodyInt)
	req, _ = http.NewRequest(http.MethodPut, "/int/second", bytes.NewReader(v))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Set the limit to 15
	bodyInt = intValue{Value: 15}
	v, _ = json.Marshal(bodyInt)
	req, _ = http.NewRequest(http.MethodPut, "/limit", bytes.NewReader(v))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Set the first string value to fizz
	bodyStr := stringValue{Value: "fizz"}
	v, _ = json.Marshal(bodyStr)
	req, _ = http.NewRequest(http.MethodPut, "/string/first", bytes.NewReader(v))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Set the second string value to fizz
	bodyStr = stringValue{Value: "buzz"}
	v, _ = json.Marshal(bodyStr)
	req, _ = http.NewRequest(http.MethodPut, "/string/second", bytes.NewReader(v))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Read the first int
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/int/first", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(formatIntvalue, int1name, 3), w.Body.String())

	// Read the second int
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/int/second", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(formatIntvalue, int2name, 5), w.Body.String())

	// Read the limit
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/limit", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(formatIntvalue, limitname, 15), w.Body.String())

	// Read the first string
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/string/first", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(formatStringvalue, string1name, "fizz"), w.Body.String())

	// Read the second string
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/string/second", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(formatStringvalue, string2name, "buzz"), w.Body.String())
}
