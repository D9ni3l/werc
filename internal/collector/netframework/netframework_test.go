//go:build windows

package netframework_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/netframework"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	// No context name required as Collector source is WMI
	testutils.FuncBenchmarkCollector(b, netframework.Name, netframework.NewWithFlags)
}
