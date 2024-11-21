//go:build windows

package hyperv_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/hyperv"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, hyperv.Name, hyperv.NewWithFlags)
}
