package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

// func TestMainHandler(t *testing.T) {

// 	assert := assert.New(t)

// 	res := httptest.NewRecorder()
// 	req := httptest.NewRequest("GET", "/", nil)

// 	mainHandler(res, req)

// 	assert.Equal(http.StatusOK, res.Code)

// 	data, _ := io.ReadAll(res.Body)
// 	assert.Equal("Hello World", string(data))
// }
