//go:build windows

package dns_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/dns"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, dns.Name, dns.NewWithFlags)
}
