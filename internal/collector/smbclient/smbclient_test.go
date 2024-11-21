//go:build windows

package smbclient_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/smbclient"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, smbclient.Name, smbclient.NewWithFlags)
}
