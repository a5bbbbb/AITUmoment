package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RedirectHandler struct {
}

func NewRedirectHandler() *RedirectHandler {
	return &RedirectHandler{}
}

func (r *RedirectHandler) Redirect(port, path string) func(c *gin.Context) {
	return func(c *gin.Context) {

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.String(400, "Cannot read request body.")
			return
		}

		url := fmt.Sprintf("%s://%s%s", "http", "localhost:"+port, path)

		fmt.Println("URL: " + url)

		proxyReq, err := http.NewRequest(c.Request.Method, url, bytes.NewReader(body))

		if err != nil {
			c.String(500, "Cannot generate new request.")
		}
		proxyReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		httpClient := http.Client{}

		resp, err := httpClient.Do(proxyReq)
		if err != nil {
			c.String(500, "Sorry service is down.")
			return
		}
		defer resp.Body.Close()

		resBody, err := io.ReadAll(resp.Body)
		if err != nil {
			c.String(500, "Cannot read response body.")
			return
		}
		c.Data(resp.StatusCode, "text/html", resBody)
	}
}
