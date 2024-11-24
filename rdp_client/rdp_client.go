package rdp_client

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"github.com/prometheus/client_golang/prometheus"
)

// Name defines the collector name for registration.
const Name = "rdp_client"

// RDPClientCollector collects metrics about active RDP clients.
type RDPClientCollector struct {
	RDPSession *prometheus.Desc
}

// NewWithFlags creates a new RDP client collector.
func NewWithFlags() *RDPClientCollector {
	return NewRDPClientCollector()
}

// NewRDPClientCollector creates a new RDPClientCollector.
func NewRDPClientCollector() *RDPClientCollector {
	return &RDPClientCollector{
		RDPSession: prometheus.NewDesc(
			"windows_rdp_client_active",
			"Information about active RDP client sessions",
			[]string{"client_name"},
			nil,
		),
	}
}

// Describe sends the metrics descriptions to the Prometheus channel.
func (c *RDPClientCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.RDPSession
}

// Collect fetches the metrics and sends them to the Prometheus channel.
func (c *RDPClientCollector) Collect(ch chan<- prometheus.Metric) {
	clientName, err := getClientName()
	if err != nil {
		fmt.Printf("Error fetching RDP client name: %v\n", err)
		return
	}

	// Report the metric with the RDP client name
	ch <- prometheus.MustNewConstMetric(
		c.RDPSession,
		prometheus.GaugeValue,
		1, // Indicating an active session
		clientName,
	)
}

// getClientName fetches the CLIENTNAME from the registry.
func getClientName() (string, error) {
	keyPath := `Volatile Environment`
	key, err := registry.OpenKey(registry.CURRENT_USER, keyPath, registry.READ)
	if err != nil {
		return "", fmt.Errorf("failed to open registry key: %v", err)
	}
	defer key.Close()

	clientName, _, err := key.GetStringValue("CLIENTNAME")
	if err == registry.ErrNotExist {
		return "No Active RDP Session", nil
	} else if err != nil {
		return "", fmt.Errorf("failed to get CLIENTNAME: %v", err)
	}

	if clientName != "" {
		return clientName, nil
	}

	return "No Active RDP Session", nil
}

