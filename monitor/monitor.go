package monitor

import (
	"net/http"
	"net/url"
	"time"

	"github.com/mfbmina/poc-prometheus-exporter/exporter"
)

type Monitor struct {
	URL         *url.URL
	SuccessCode int
	PExporter   *exporter.PExporter
	PGExporter  *exporter.PGExporter
}

func (m Monitor) Start() {
	for {
		time.Sleep(1 * time.Second)
		err := doRequest(m.URL.String(), m.SuccessCode)
		if err != nil {
			m.PExporter.RecordFailure()
			m.PGExporter.RecordFailure()
			continue
		}

		m.PExporter.RecordSuccess()
		m.PGExporter.RecordSuccess()
	}
}

func doRequest(u string, expectedStatus int) error {
	resp, err := http.Get(u)
	if err != nil {
		return err
	}

	if resp.StatusCode != expectedStatus {
		return err
	}

	return nil
}
