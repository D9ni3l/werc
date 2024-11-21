//go:build windows

package update_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/update"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, "printer", update.NewWithFlags)
}
