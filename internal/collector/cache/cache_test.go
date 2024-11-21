//go:build windows

package cache_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/cache"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, cache.Name, cache.NewWithFlags)
}
