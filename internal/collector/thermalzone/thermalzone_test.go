//go:build windows

package thermalzone_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/thermalzone"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, thermalzone.Name, thermalzone.NewWithFlags)
}
