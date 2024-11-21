//go:build windows

package tcp_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/tcp"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, tcp.Name, tcp.NewWithFlags)
}

func TestCollector(t *testing.T) {
	testutils.TestCollector(t, tcp.New, nil)
}
