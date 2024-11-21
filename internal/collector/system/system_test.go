//go:build windows

package system_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/system"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, system.Name, system.NewWithFlags)
}
