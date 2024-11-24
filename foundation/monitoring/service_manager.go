package monitoring

import (
	"sync"
)

// Service manager that tracks the active services
var (
	servicesMetrics = make(map[string]bool) // Active services
	mu              sync.Mutex              // Mutex to guard servicesMetrics
)

// Start and stop service-specific metrics collection
func StartMetricsForService(serviceName string) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := servicesMetrics[serviceName]; !exists {
		// Register the service for metrics collection
		registerService(serviceName)
	}
}

func StopMetricsForService(serviceName string) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := servicesMetrics[serviceName]; exists {
		// Unregister the service from metrics collection
		unregisterService(serviceName)
	}
}
