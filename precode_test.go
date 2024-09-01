package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// проверка на непустое тело и статус ОК
func TestBodyNotEmpty(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4?city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body
	require.Equal(t, http.StatusOK, responseRecorder.Code) //Статус ОК
	assert.NotEmpty(t, body)                               //непустое тело
}

// не тот город
func TestWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4?city=jopa", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	body := responseRecorder.Body
	assert.NotEmpty(t, body)
}

// Счётчик больше максимума
func TestCounterMoreThanMax(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=100?city=jopa", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Len(t, strings.Split(body, ","), 100)
}
