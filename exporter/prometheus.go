package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type PExporter struct {
	successRate prometheus.Counter
	failureRate prometheus.Counter
}

func NewPrometheusExporter() *PExporter {
	successRate := promauto.NewCounter(prometheus.CounterOpts{
		Name: "success_rate",
		Help: "The total number of succeded events",
	})

	failureRate := promauto.NewCounter(prometheus.CounterOpts{
		Name: "failure_rate",
		Help: "The total number of failed events",
	})

	return &PExporter{
		successRate: successRate,
		failureRate: failureRate,
	}
}

func (p PExporter) RecordSuccess() {
	p.successRate.Inc()
}

func (p PExporter) RecordFailure() {
	p.failureRate.Inc()
}
