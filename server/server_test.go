package server_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRunServer(t *testing.T) {
	e := echo.New()

	ts := httptest.NewServer(e)
	defer ts.Close()

	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is running")
	})

	res, err := http.Get(ts.URL + "/test")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	assert.NoError(t, err)

	assert.Equal(t, "Server is running", string(body))
}
