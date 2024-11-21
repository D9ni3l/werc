//go:build windows

package smtp_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/smtp"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, smtp.Name, smtp.NewWithFlags)
}
