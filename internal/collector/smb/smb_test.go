//go:build windows

package smb_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/smb"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, smb.Name, smb.NewWithFlags)
}
