//go:build windows

package physical_disk_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/physical_disk"
	"github.com/D9ni3l/werc/internal/types"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, physical_disk.Name, physical_disk.NewWithFlags)
}

func TestCollector(t *testing.T) {
	testutils.TestCollector(t, physical_disk.New, &physical_disk.Config{
		DiskInclude: types.RegExpAny,
	})
}
