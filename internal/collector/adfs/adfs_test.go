//go:build windows

package adfs_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/adfs"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, adfs.Name, adfs.NewWithFlags)
}
