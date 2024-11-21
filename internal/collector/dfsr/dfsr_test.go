//go:build windows

package dfsr_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/dfsr"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, dfsr.Name, dfsr.NewWithFlags)
}
