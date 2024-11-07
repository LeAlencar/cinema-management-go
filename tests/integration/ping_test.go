package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"cinema-project-go/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	return r
}

func TestPingIntegration(t *testing.T) {
	// Setup router
	router := setupRouter()

	// Create test server
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Make HTTP request
	resp, err := http.Get(ts.URL + "/ping")
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Assert status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Parse response
	var response handlers.PingResponse
	err = json.NewDecoder(resp.Body).Decode(&response)

	// Assert response
	assert.NoError(t, err)
	assert.Equal(t, "pong", response.Message)
	assert.NotEmpty(t, response.Timestamp)
	assert.True(t, response.Timestamp.Before(time.Now()))
}
