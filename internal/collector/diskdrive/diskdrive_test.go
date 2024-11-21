//go:build windows

package diskdrive_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/diskdrive"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, diskdrive.Name, diskdrive.NewWithFlags)
}
