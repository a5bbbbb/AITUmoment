package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type ApiGatewayMetrics struct {
	requestCount   prometheus.Counter
	requestLatency *prometheus.HistogramVec
}

func NewApiGatewayMetrics() *ApiGatewayMetrics {
	return &ApiGatewayMetrics{
		requestCount: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "http_requests_counter",
				Help: "Total number of HTTP requests received",
			},
		),
		requestLatency: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_ms",
				Help:    "Histogram of the duration of HTTP requests processed.",
				Buckets: prometheus.LinearBuckets(0, 50, 9),
			},
			[]string{"method", "path"},
		),
	}
}

func (a *ApiGatewayMetrics) RegisterRoutes() {
	prometheus.MustRegister(a.requestCount)
	prometheus.MustRegister(a.requestLatency)
}

func (a *ApiGatewayMetrics) HttpRequestTotal(c *gin.Context) {
	a.requestCount.Inc()
	c.Next()
}

func (a *ApiGatewayMetrics) HttpRequestLatency(c *gin.Context) {
	start := time.Now()
	method := c.Request.Method
	path := c.Request.URL.Path
	c.Next()
	duration := time.Since(start)
	a.requestLatency.With(
		prometheus.Labels{
			"method": method,
			"path":   path,
		},
	).Observe(float64(duration.Milliseconds()))
}
