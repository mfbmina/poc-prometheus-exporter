package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const serviceURL = "http://localhost:1313"
const successCode = 200

func main() {
	// u, err := url.Parse(serviceURL)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// m := monitor.Monitor{
	// 	SuccessCode: successCode,
	// 	URL:         u,
	// 	PExporter:   exporter.NewPrometheusExporter(),
	// 	PGExporter:  exporter.NewPushGatewayExporter(),
	// }

	// go m.Start()

	fmt.Println("Monitoring service...")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
