package server

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/p-pawel/go-challenge/database"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {

	_, _, _ = sqlmock.NewWithDSN("sqlmock_db_0")
	database.DB, _ = gorm.Open("sqlmock", "sqlmock_db_0")

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

	t.Run("Should post booking (no validation so far)", func(t *testing.T) {

		// given
		request, _ := http.NewRequest(http.MethodPost, "/booking", strings.NewReader("{ }"))
		response := httptest.NewRecorder()

		// when
		router.ServeHTTP(response, request)
		log.Println(response)

		// then
		assert.Equal(t, http.StatusOK, response.Code)
		assert.NotNil(t, response.Body.String())
	})

}
