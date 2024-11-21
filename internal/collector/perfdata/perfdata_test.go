//go:build windows

package perfdata_test

import (
	"testing"

	"github.com/alecthomas/kingpin/v2"
	"github.com/D9ni3l/werc/internal/collector/perfdata"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	perfDataObjects := `[{"object":"Processor Information","instances":["*"],"counters":{"*": {}}}]`
	kingpin.CommandLine.GetArg("collector.perfdata.objects").StringVar(&perfDataObjects)

	testutils.FuncBenchmarkCollector(b, perfdata.Name, perfdata.NewWithFlags)
}
