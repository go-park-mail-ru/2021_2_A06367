package middleware

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"net/http"
	"time"
)

const (
	ServiceAuthLabel  = "auth"
	ServiceMainLabel  = "main"
	ServiceFilmsLabel = "films"
)

const (
	ServiceName = "service"
	StartTime   = "start"
	FullTime    = "duration"
	URL         = "url"
	Method      = "method"
	StatusCode  = "code"
)

type writer struct {
	http.ResponseWriter
	statusCode int
}

func NewWriter(w http.ResponseWriter) *writer {
	// WriteHeader(int) is not called if our response implicitly returns 200 OK, so
	// we default to that status code.
	return &writer{w, http.StatusOK}
}

func (w *writer) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

type MetricsMiddleware struct {
	metric *prometheus.GaugeVec
	name   string
}

func NewMetricsMiddleware() *MetricsMiddleware {
	return &MetricsMiddleware{}
}

func (m *MetricsMiddleware) Register(name string) {

	m.name = name
	gauge := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: name,
			Help: fmt.Sprintf("SLO for service %s", name),
		},
		[]string{
			ServiceName, StartTime, FullTime, URL, Method, StatusCode,
		})

	m.metric = gauge

	rand.Seed(time.Now().Unix())
	prometheus.MustRegister(gauge)
}

func (m *MetricsMiddleware) LogMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqId := fmt.Sprintf("%016x", rand.Int())[:10]
		ctx = context.WithValue(ctx, "ReqId", reqId)

		start := time.Now()

		wrapper := NewWriter(w)

		next.ServeHTTP(wrapper, r.WithContext(ctx))

		m.metric.With(prometheus.Labels{
			ServiceName: m.name,
			StartTime:   start.String(),
			FullTime:    time.Since(start).String(),
			URL:         r.URL.Path,
			Method:      r.Method,
			StatusCode:  fmt.Sprintf("%d", wrapper.statusCode),
		}).Inc()

	})
}
