//go:build windows

package scheduled_task_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/scheduled_task"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, scheduled_task.Name, scheduled_task.NewWithFlags)
}
