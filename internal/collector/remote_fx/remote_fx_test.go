//go:build windows

package remote_fx_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/remote_fx"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, remote_fx.Name, remote_fx.NewWithFlags)
}
