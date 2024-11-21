//go:build windows

package exchange_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/exchange"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, exchange.Name, exchange.NewWithFlags)
}
