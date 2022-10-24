package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPutIntFirst(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	body := intValue{Value: 10}
	v, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPut, "/int/first", bytes.NewReader(v))

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"value":10}`, w.Body.String())
}
