//go:build windows

package mssql_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/mssql"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func BenchmarkCollector(b *testing.B) {
	testutils.FuncBenchmarkCollector(b, mssql.Name, mssql.NewWithFlags)
}
