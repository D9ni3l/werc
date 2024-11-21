//go:build windows

package filetime_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/filetime"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, filetime.Name, filetime.NewWithFlags)
}

func TestCollector(t *testing.T) {
	testutils.TestCollector(t, filetime.New, &filetime.Config{
		FilePatterns: []string{"*.*"},
	})
}
