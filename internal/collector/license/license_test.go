//go:build windows

package license_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/license"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, license.Name, license.NewWithFlags)
}

func TestCollector(t *testing.T) {
	testutils.TestCollector(t, license.New, nil)
}
