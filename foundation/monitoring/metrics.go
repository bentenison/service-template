package monitoring

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// Prometheus metrics for HTTP request tracking
var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests processed",
		},
		[]string{"service", "method", "status"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"service", "method", "status"},
	)

	// System health metrics
	cpuUsage           = prometheus.NewGauge(prometheus.GaugeOpts{Name: "system_cpu_usage_percentage", Help: "System CPU usage in percentage"})
	memUsage           = prometheus.NewGauge(prometheus.GaugeOpts{Name: "system_memory_usage_bytes", Help: "System memory usage in bytes"})
	diskUsage          = prometheus.NewGauge(prometheus.GaugeOpts{Name: "system_disk_usage_percentage", Help: "System disk usage in percentage"})
	registeredServices = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "registered_services_count",
			Help: "Number of registered services",
		},
		[]string{"service", "status"},
	)
)

func init() {
	// once.Do(func() {
	log.Println("Initializing Prometheus metrics...")
	prometheus.Register(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(memUsage)
	prometheus.MustRegister(diskUsage)
	prometheus.MustRegister(registeredServices)
	promhttp.Handler()
	// Register the metrics with Prometheus

}

// Collect system-level metrics such as CPU, memory, and disk usage
func CollectSystemMetrics() {
	go func() {
		for {
			// Collect CPU usage
			cpuPercent, err := getCPUUsage()
			if err != nil {
				log.Printf("Error collecting CPU usage: %v", err)
			} else {
				cpuUsage.Set(cpuPercent)
			}
			// log.Println("CPU", cpuPercent)
			memPercent, err := getMemoryUsage()
			if err != nil {
				log.Printf("Error collecting memory usage: %v", err)
			} else {
				memUsage.Set(memPercent)
			}

			diskPercent, err := getDiskUsage()
			if err != nil {
				log.Printf("Error collecting disk usage: %v", err)
			} else {
				diskUsage.Set(diskPercent)
			}
			time.Sleep(5 * time.Second)
		}
	}()
}

// getCPUUsage gets the CPU usage percentage
func getCPUUsage() (float64, error) {
	// Get the overall CPU usage (average for all cores)
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}

	// Return the first value in the percentages slice as the overall CPU usage
	return percentages[0], nil
}

// getMemoryUsage gets the memory usage percentage
func getMemoryUsage() (float64, error) {
	// Get the virtual memory stats
	vmem, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}

	// Return the percentage of memory usage
	return vmem.UsedPercent, nil
}

// getDiskUsage gets the disk usage percentage of the root partition
func getDiskUsage() (float64, error) {
	// Get disk usage stats for the root directory
	usage, err := disk.Usage("/")
	if err != nil {
		return 0, err
	}

	// Return the percentage of disk usage
	return usage.UsedPercent, nil
}

// Register a service to start collecting its metrics
func registerService(serviceName string) {
	registeredServices.WithLabelValues(serviceName, "active").Set(1)
}

// Unregister a service to stop collecting its metrics
func unregisterService(serviceName string) {
	registeredServices.WithLabelValues(serviceName, "inactive").Set(0)
}

// Collect HTTP request metrics for a service
func CollectHTTPRequestMetrics(serviceName, method, status string, duration float64) {
	httpRequestsTotal.WithLabelValues(serviceName, method, status).Inc()
	httpRequestDuration.WithLabelValues(serviceName, method, status).Observe(duration)
}

// Expose /metrics endpoint for Prometheus scraping
func ExposeMetrics(addr string) {
	http.Handle("/metrics", promhttp.Handler())
	// go func() {
	log.Printf("Metrics server started at %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("Error starting HTTP server: ", err)
	}
	// }()
}
