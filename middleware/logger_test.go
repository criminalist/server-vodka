package middleware

import (
	"errors"
	"github.com/insionng/vodka"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogger(t *testing.T) {
	e := vodka.New()
	req, _ := http.NewRequest(vodka.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := vodka.NewContext(req, vodka.NewResponse(rec), e)

	// Status 2xx
	h := func(c *vodka.Context) error {
		return c.String(http.StatusOK, "test")
	}
	Logger()(h)(c)

	// Status 3xx
	rec = httptest.NewRecorder()
	c = vodka.NewContext(req, vodka.NewResponse(rec), e)
	h = func(c *vodka.Context) error {
		return c.String(http.StatusTemporaryRedirect, "test")
	}
	Logger()(h)(c)

	// Status 4xx
	rec = httptest.NewRecorder()
	c = vodka.NewContext(req, vodka.NewResponse(rec), e)
	h = func(c *vodka.Context) error {
		return c.String(http.StatusNotFound, "test")
	}
	Logger()(h)(c)

	// Status 5xx with empty path
	req, _ = http.NewRequest(vodka.GET, "", nil)
	rec = httptest.NewRecorder()
	c = vodka.NewContext(req, vodka.NewResponse(rec), e)
	h = func(c *vodka.Context) error {
		return errors.New("error")
	}
	Logger()(h)(c)
}
