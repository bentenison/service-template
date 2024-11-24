// package main

// import (
// 	"log"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	"github.com/bentenison/microservice/api/cmd/service/metrics/monitoring"
// 	// Replace with your actual package import path
// )

// func main() {
// 	// Start collecting system health metrics (CPU, memory, disk)
// 	// go monitoring.collectSystemMetrics()

// 	// Expose the /metrics endpoint for Prometheus scraping
// 	monitoring.ExposeMetrics(":8007")

// 	// Gracefully shut down when receiving a termination signal
// 	go gracefulShutdown()

// 	// Start the main server loop
// 	// select {}
// }

// // gracefulShutdown handles graceful termination of the monitoring service
// func gracefulShutdown() {
// 	// Channel to receive OS signals for graceful shutdown
// 	stop := make(chan os.Signal, 1)
// 	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

// 	// Block until a shutdown signal is received
// 	<-stop

// 	log.Println("Received shutdown signal, shutting down the monitoring service...")

// 	// Perform any necessary cleanup here (e.g., stopping background metrics collection)
// 	time.Sleep(2 * time.Second) // Simulating cleanup time
// 	log.Println("Monitoring service shutdown complete.")
// }

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/expfmt"
)

// CollectorService structure
type CollectorService struct {
	targetURLs []Targets
	metrics    *prometheus.GaugeVec
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// NewCollectorService initializes a new CollectorService
func NewCollectorService(targetURLs []Targets) *CollectorService {
	return &CollectorService{
		targetURLs: targetURLs,
		metrics: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "external_service_metrics_collected",
				Help: "Metrics collected from external services",
			},
			[]string{"service_name", "metric_name"},
		),
	}
}

// ScrapeMetrics collects metrics from multiple services and adds them to the GaugeVec
func (c *CollectorService) ScrapeMetrics() {
	for _, targetURL := range c.targetURLs {
		// Scrape the /metrics endpoint from the target service
		resp, err := http.Get(targetURL.Url + "/metrics")
		if err != nil {
			log.Printf("Error scraping %s: %v", targetURL, err)
			continue
		}
		defer resp.Body.Close()

		// Read the response body
		// body, err := io.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Printf("Error reading body of %s: %v", targetURL, err)
		// 	continue
		// }

		// // Log the raw body for debugging
		// log.Printf("Metrics from %s:\n%s", targetURL, body)

		// Parse the metrics from the response body
		parser := expfmt.TextParser{}
		metricFamilies, err := parser.TextToMetricFamilies(resp.Body)
		if err != nil {
			log.Printf("Error parsing metrics from %s: %v", targetURL, err)
			continue
		}

		// Process the parsed metrics and store them
		for _, mf := range metricFamilies {
			for _, m := range mf.GetMetric() {
				// Add each metric to the centralized metrics collector
				c.metrics.WithLabelValues(targetURL.Name, mf.GetName()).Set(m.GetGauge().GetValue())
			}
		}
	}
}

// ExposeMetrics exposes the aggregated metrics through a /metrics endpoint
func (c *CollectorService) ExposeMetrics() {
	// Register the GaugeVec with Prometheus
	prometheus.MustRegister(c.metrics)

	// Expose metrics to Prometheus via HTTP
	http.Handle("/metrics", promhttp.Handler())

	// Start the HTTP server to expose the /metrics endpoint
	log.Println("Centralized Metrics Collector is running on :8007...")
	if err := http.ListenAndServe(":8007", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

type Targets struct {
	Url  string
	Name string
}

func main() {
	// List of target service URLs to collect metrics from
	targetURLs := []Targets{{Url: "http://localhost:8004", Name: "executor-service"}, {Url: "http://localhost:8003", Name: "broker-service"}}

	// Create a new collector service
	collector := NewCollectorService(targetURLs)

	// Start a separate goroutine to periodically scrape metrics from the target services
	go func() {
		for {
			// Collect metrics every 10 seconds
			collector.ScrapeMetrics()
			time.Sleep(10 * time.Second)
		}
	}()

	// Expose the collected metrics on the /metrics endpoint
	collector.ExposeMetrics()
}
