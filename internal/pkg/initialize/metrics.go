package initialize

import (
	"go-web-cli/internal/pkg/metrics"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

type MetricConfig struct {
	Open       bool
	Prometheus struct {
		Namespace string
		SubSystem string
		Url       string
		Bind      string
	}
}

func Metric() error {
	metricCfg := MetricConfig{}

	err := viper.UnmarshalKey("metrics", &metricCfg)
	if err != nil {
		return err
	}

	if !metricCfg.Open {
		return nil
	}

	metrics.InitMetrics(metricCfg.Prometheus.Namespace, metricCfg.Prometheus.SubSystem)

	go func() {
		http.Handle(metricCfg.Prometheus.Url, promhttp.Handler())

		server := &http.Server{
			Addr:              metricCfg.Prometheus.Bind,
			ReadHeaderTimeout: 10 * time.Second,
		}

		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	return nil
}
