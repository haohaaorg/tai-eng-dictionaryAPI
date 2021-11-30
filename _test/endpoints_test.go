package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/noernova/tai-eng-dictionaryAPI/controllers"
	"github.com/stretchr/testify/assert"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/", controllers.Index)
	router.GET("/api/v1/api_key=NO8p3FC4qMrTzx1RUjRXNXWrqlLa8DkDjmRgt7s9rDE=/eng/:text", controllers.Eng_to_Shn)
	router.GET("/api/v1/api_key=NO8p3FC4qMrTzx1RUjRXNXWrqlLa8DkDjmRgt7s9rDE=/shn/:text", controllers.Shn_to_Eng)
	return router
}
func TestAllEndPoint(t *testing.T) {
	r := Router()
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestIndexEndPoint(t *testing.T) {
	r := Router()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)

	var resp map[string]interface{}
	available_endpoints := []interface {}([]interface {}{"/api/v1/apikey={apikey}eng/:text", "/api/v1/apikey={apikey}shn/:text"})

	body := w.Body.String()
	var resp_body = []byte(body)

	err := json.Unmarshal(resp_body, &resp)

	assert.Nil(t, err)
	assert.Equal(t, "1.0.0", resp["version"])
	assert.Equal(t, "Welcome to shn_eng-dic API.", resp["message"])
	assert.Equal(t, available_endpoints , resp["available_endpoints"])
}
