package exporter

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/push"
)

type PGExporter struct {
	successRate prometheus.Counter
	failureRate prometheus.Counter
}

func NewPushGatewayExporter() *PGExporter {
	successRate := promauto.NewCounter(prometheus.CounterOpts{
		Name: "success_rate_pg",
		Help: "The total number of succeded events",
	})

	failureRate := promauto.NewCounter(prometheus.CounterOpts{
		Name: "failure_rate_pg",
		Help: "The total number of failed events",
	})

	return &PGExporter{
		successRate: successRate,
		failureRate: failureRate,
	}
}

func (p PGExporter) RecordSuccess() {
	p.successRate.Inc()

	err := push.New("http://pushgateway:9091", "pg").
		Collector(p.successRate).
		Grouping("myapp", "success_rate_pg").
		Push()

	if err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}

func (p PGExporter) RecordFailure() {
	p.failureRate.Inc()

	err := push.New("http://pushgateway:9091", "pg").
		Collector(p.failureRate).
		Grouping("myapp", "failure_rate_pg").
		Push()

	if err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}
