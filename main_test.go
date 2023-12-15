package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCorrectRequest(t *testing.T) {
	//GIVEN
	count := 2
	city := "moscow"
	url := fmt.Sprintf("/cafe?count=%d&city=%s", count, city)
	req := httptest.NewRequest(http.MethodGet, url, nil)

	expectedBody := strings.Join(cafeList[city][:2], ",")

	//WHEN
	responseRecorder := serveHttpRequest(req)
	actualBody := string(responseRecorder.Body.Bytes()[:])

	//THEN
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, expectedBody, actualBody)

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	//GIVEN
	count := 5
	city := "moscow"
	url := fmt.Sprintf("/cafe?count=%d&city=%s", count, city)
	req := httptest.NewRequest(http.MethodGet, url, nil)

	//expectedBody := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	expectedBody := strings.Join(cafeList[city][0:4], ",")

	//WHEN
	responseRecorder := serveHttpRequest(req)
	actualBody := string(responseRecorder.Body.Bytes()[:])

	//THEN
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, expectedBody, actualBody)

}

func TestMainHandlerWhenUnsupportedCity(t *testing.T) {
	//GIVEN
	count := 2
	city := "nsk"
	url := fmt.Sprintf("/cafe?count=%d&city=%s", count, city)
	req := httptest.NewRequest(http.MethodGet, url, nil)

	expectedBody := "wrong city value"

	//WHEN
	responseRecorder := serveHttpRequest(req)
	actualBody := string(responseRecorder.Body.Bytes()[:])

	//THEN
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, expectedBody, actualBody)

}

func serveHttpRequest(request *http.Request) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, request)
	return responseRecorder
}
