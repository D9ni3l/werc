//go:build windows

package terminal_services_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/terminal_services"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, terminal_services.Name, terminal_services.NewWithFlags)
}
