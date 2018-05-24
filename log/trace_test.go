package log

import (
	"testing"
)

func TestTraceRace(t *testing.T) {
	traceBegin("1", "hello", "my", "husband")
	go traceBegin("2", "hello", "my", "wift")
}
