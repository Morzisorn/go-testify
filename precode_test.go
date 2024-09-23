package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOK(t *testing.T) {
	count := 2
	req := httptest.NewRequest("Get", fmt.Sprintf("/cafe?count=%d&city=moscow", count), nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotNil(t, responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("Get", "/cafe?count=6&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	cafes := strings.Split(responseRecorder.Body.String(), ",")

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, totalCount, len(cafes))
}

func TestMainHandlerWhenUnsupportedCity(t *testing.T) {
	city := "Tula"
	respBody := "wrong city value"
	req := httptest.NewRequest("Get", fmt.Sprintf("/cafe?count=3&city=%s", city), nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, respBody, responseRecorder.Body.String())
}
