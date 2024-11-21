//go:build windows

package memory_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/memory"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, memory.Name, memory.NewWithFlags)
}

func TestCollector(t *testing.T) {
	testutils.TestCollector(t, memory.New, nil)
}
