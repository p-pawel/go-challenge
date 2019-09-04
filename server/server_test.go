package server

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestServer(t *testing.T) {

	server := &RocketServer{}
	t.Run("Server should be available", func(t *testing.T) {

		// given
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "/", nil)

		// when
		server.ServeHTTP(response, request)

		// then
		assert.Equal(t, http.StatusOK, response.Code)
	})

	
}