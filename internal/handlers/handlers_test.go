package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/ping", Ping)
	return r
}

func TestPing(t *testing.T) {
	t.Run("returns 200 status code", func(t *testing.T) {
		// Arrange
		router := setupTestRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/ping", nil)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("returns correct message", func(t *testing.T) {
		// Arrange
		router := setupTestRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/ping", nil)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		var response PingResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "pong", response.Message)
	})

	t.Run("returns valid timestamp", func(t *testing.T) {
		// Arrange
		router := setupTestRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
		beforeTest := time.Now()

		// Act
		router.ServeHTTP(w, req)

		// Assert
		var response PingResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.True(t, response.Timestamp.After(beforeTest.Add(-time.Second)))
		assert.True(t, response.Timestamp.Before(time.Now().Add(time.Second)))
	})

	t.Run("returns correct content type", func(t *testing.T) {
		// Arrange
		router := setupTestRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/ping", nil)

		// Act
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	})

	t.Run("handles multiple consecutive requests", func(t *testing.T) {
		// Arrange
		router := setupTestRouter()

		for i := 0; i < 3; i++ {
			t.Run(fmt.Sprintf("request_%d", i), func(t *testing.T) {
				w := httptest.NewRecorder()
				req, _ := http.NewRequest(http.MethodGet, "/ping", nil)

				// Act
				router.ServeHTTP(w, req)

				// Assert
				assert.Equal(t, http.StatusOK, w.Code)
				var response PingResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				require.NoError(t, err)
				assert.Equal(t, "pong", response.Message)
			})
		}
	})
}
