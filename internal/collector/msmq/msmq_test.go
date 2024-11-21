//go:build windows

package msmq_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/msmq"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	// No context name required as Collector source is WMI
	testutils.FuncBenchmarkCollector(b, msmq.Name, msmq.NewWithFlags)
}
