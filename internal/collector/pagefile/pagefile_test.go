//go:build windows

package pagefile_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/pagefile"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, pagefile.Name, pagefile.NewWithFlags)
}

func TestCollector(t *testing.T) {
	testutils.TestCollector(t, pagefile.New, nil)
}
