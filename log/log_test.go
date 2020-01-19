package log

import "testing"

func TestPrint(t *testing.T) {
	var _log Logger = `Concurrency`
	_log.Print(`testing log`)
}
