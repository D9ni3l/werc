//go:build windows

package adcs_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/adcs"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, adcs.Name, adcs.NewWithFlags)
}
