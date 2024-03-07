package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
)

var (
	RequestLatency *prometheus.HistogramVec
)

func InitMetrics(namespace, subsystem string) {
	var buckets []float64
	err := viper.UnmarshalKey("metrics.prometheus.buckets", &buckets)
	if err != nil {
		buckets = []float64{0, 1, 2.5, 5, 10, 25, 50, 100, 250, 500}
	}

	// metric request count and latency.
	RequestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "http_request_operation_duration_milliseconds",
			Help:      "Count and latency of operation request in millisecond.",
			Buckets:   buckets, // ms
		},
		[]string{"host", "method", "path", "status"},
	)

	// must register counter on init
	prometheus.MustRegister(RequestLatency)
}
