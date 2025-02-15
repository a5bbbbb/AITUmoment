package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RedirectHandler struct {
	httpClient http.Client
}

func NewRedirectHandler() *RedirectHandler {
	return &RedirectHandler{
		httpClient: http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

func (r *RedirectHandler) Redirect(host, port, path string) func(c *gin.Context) {
	return func(c *gin.Context) {

		// Form the target url
		url := fmt.Sprintf("%s://%s%s", "http", host+":"+port, c.Request.RequestURI)

		// Get orig request body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.String(400, "Cannot read request body.")
			return
		}

		// Create proxy request with body of the received request
		proxyReq, err := http.NewRequest(c.Request.Method, url, bytes.NewReader(body))

		if err != nil {
			c.String(500, "Cannot generate proxy request.")
		}

		// Copy headers
		proxyReq.Header = c.Request.Header

		// Copy cookies into the proxy request
		for _, value := range c.Request.Cookies() {
			proxyReq.AddCookie(value)
		}

		// Send the proxy request
		resp, err := r.httpClient.Do(proxyReq)
		if err != nil {
			c.String(500, "Sorry service is down.")
			return
		}

		// Need to close body reader
		defer resp.Body.Close()

		// Read the response body
		resBody, err := io.ReadAll(resp.Body)
		if err != nil {
			c.String(500, "Cannot read response body.")
			return
		}

		// Copy response headers into the response
		for key, values := range resp.Header {

			for _, one := range values {
				c.Header(key, one)
			}
		}

		c.Data(resp.StatusCode, resp.Header.Get("Content-type"), resBody)
	}
}
