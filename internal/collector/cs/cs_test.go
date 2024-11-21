//go:build windows

package cs_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/cs"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, cs.Name, cs.NewWithFlags)
}
