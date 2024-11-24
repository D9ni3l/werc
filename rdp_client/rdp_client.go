package rdp_client

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"github.com/prometheus/client_golang/prometheus"
	"log/slog"
)

// Name defines the collector name for registration.
const Name = "rdp_client"

// Config defines the configuration for the RDP client collector.
type Config struct {
	Enabled bool `yaml:"enabled"`
}

// ConfigDefaults provides the default configuration for the RDP client collector.
var ConfigDefaults = Config{
	Enabled: true,
}

// RDPClientCollector collects metrics about active RDP clients.
type RDPClientCollector struct {
	RDPSession *prometheus.Desc
}

// NewWithFlags creates a new RDPClientCollector with flags.
func NewWithFlags(app *kingpin.Application) *RDPClientCollector {
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

// Build satisfies the Collector interface and accepts a session.
func (c *RDPClientCollector) Build(logger *slog.Logger, session *mi.Session) error {
	logger.Info("RDPClientCollector Build invoked")
	// Log that the session is unused
	logger.Info("Session is not used in this collector")
	return nil
}

// getClientName fetches the CLIENTNAME from the registry.
func getClientName() (string, error) {
	// Open the registry key for Volatile Environment
	keyPath := `Volatile Environment`
	key, err := registry.OpenKey(registry.CURRENT_USER, keyPath, registry.READ)
	if err != nil {
		return "", fmt.Errorf("failed to open registry key: %v", err)
	}
	defer key.Close()

	// Attempt to get the CLIENTNAME value
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

