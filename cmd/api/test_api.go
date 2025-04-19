package main

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMainExecution(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Подменяем зависимости на моки
	os.Setenv("USE_IN_MEMORY", "true")
	os.Setenv("PORT", "8080")

	t.Run("Health check", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		go func() {
			main()
		}()
		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
