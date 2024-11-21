//go:build windows

package logical_disk_test

import (
	"testing"

	"github.com/alecthomas/kingpin/v2"
	"github.com/D9ni3l/werc/internal/collector/logical_disk"
	"github.com/D9ni3l/werc/internal/types"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	// Whitelist is not set in testing context (kingpin flags not parsed), causing the Collector to skip all disks.
	localVolumeInclude := ".+"
	kingpin.CommandLine.GetArg("collector.logical_disk.volume-include").StringVar(&localVolumeInclude)
	testutils.FuncBenchmarkCollector(b, "logical_disk", logical_disk.NewWithFlags)
}

func TestCollector(t *testing.T) {
	testutils.TestCollector(t, logical_disk.New, &logical_disk.Config{
		VolumeInclude: types.RegExpAny,
	})
}
