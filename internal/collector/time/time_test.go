//go:build windows

package time_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/time"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, time.Name, time.NewWithFlags)
}

func TestCollector(t *testing.T) {
	testutils.TestCollector(t, time.New, nil)
}
