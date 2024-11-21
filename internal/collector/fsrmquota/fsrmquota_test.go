//go:build windows

package fsrmquota_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/fsrmquota"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, fsrmquota.Name, fsrmquota.NewWithFlags)
}
