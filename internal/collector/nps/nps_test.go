//go:build windows

package nps_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/nps"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, nps.Name, nps.NewWithFlags)
}
