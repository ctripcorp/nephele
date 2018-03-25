package output

import (
	"testing"
)

func TestDump(t *testing.T) {
	dc := &DumpConfig{
		"info",
		"/Users/CINTS/nephele/log",
		1,
	}
	o, err := dc.Build()
	if err != nil {
		t.Error(err)
	}
	for {
		o.Write([]byte("h1z1o4f3"), "info")
	}
}
