//go:build windows

package ad_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/ad"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, ad.Name, ad.NewWithFlags)
}
