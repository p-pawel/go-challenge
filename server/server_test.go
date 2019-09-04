package server

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {

	router := SetupRouter()
	t.Run("Server should be available", func(t *testing.T) {

		// given
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "/", nil)

		// when
		router.ServeHTTP(response, request)

		// then
		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should get bookings (empty array as none created)", func(t *testing.T) {

		// given
		request, _ := http.NewRequest(http.MethodGet, "/booking", nil)
		response := httptest.NewRecorder()

		// when
		router.ServeHTTP(response, request)

		// then
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "[]\n", response.Body.String())
	})

}
