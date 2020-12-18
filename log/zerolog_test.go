package log

import "testing"

func TestLog(t *testing.T) {
	PrepareAll("dxc", ".")
	Info().Msg("")
}
