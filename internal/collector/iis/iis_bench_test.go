//go:build windows

package iis_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/iis"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, iis.Name, iis.NewWithFlags)
}
