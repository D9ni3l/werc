//go:build windows

package dhcp_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/dhcp"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, dhcp.Name, dhcp.NewWithFlags)
}
