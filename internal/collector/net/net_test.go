//go:build windows

package net_test

import (
	"testing"

	"github.com/D9ni3l/werc/internal/collector/net"
	"github.com/D9ni3l/werc/internal/utils/testutils"
)

func TestCollector(t *testing.T) {
	testutils.TestCollector(t, net.New, nil)
}
