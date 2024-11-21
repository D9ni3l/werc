package rdp_client

import (
    "testing"

    "github.com/prometheus/client_golang/prometheus/testutil"
)

func TestRDPClientCollector(t *testing.T) {
    collector := NewRDPClientCollector()
    if err := testutil.CollectAndCompare(collector, nil); err != nil {
        t.Errorf("unexpected collecting result: %s", err)
    }
}

