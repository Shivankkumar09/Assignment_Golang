package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/Shivankkumar09/Assignment_Golang/controllers"
)


func TestLoginSuccess(t *testing.T) {



	// Setup Gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/login", controllers.Login)

	// Mock input
	jsonBody := []byte(`{"username":"testuser1","password":"12345678"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}
