//go:build windows

package container_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/container"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, container.Name, container.NewWithFlags)
}
