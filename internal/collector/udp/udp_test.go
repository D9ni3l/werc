//go:build windows

package udp_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/udp"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, udp.Name, udp.NewWithFlags)
}

func TestCollector(t *testing.T) {
	testutils.TestCollector(t, udp.New, nil)
}
