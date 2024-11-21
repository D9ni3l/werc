//go:build windows

package vmware_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/vmware"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, vmware.Name, vmware.NewWithFlags)
}
