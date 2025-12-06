package proxy

import (
	"net/http/httputil"
	"net/url"

	"github.com/labstack/echo/v4"
)

func NewReverseProxy(target string) echo.HandlerFunc {
	t, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(t)

	return func(c echo.Context) error {
		req := c.Request()

		req.URL.Scheme = t.Scheme
		req.URL.Host = t.Host
		req.Host = t.Host

		proxy.ServeHTTP(c.Response(), req)
		return nil
	}
}
