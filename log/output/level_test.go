package output

import (
	"testing"
)

func TestLevel(t *testing.T) {
	t2s := []struct {
		desc string
		f    func() int
		want int
	}{
		{"debug == 4", func() int { return levelInt("debug") }, 4},
		{"info == 3", func() int { return levelInt("info") }, 3},
		{"warning == 2", func() int { return levelInt("warning") }, 2},
		{"error == 1", func() int { return levelInt("error") }, 1},
		{"fatal == 0", func() int { return levelInt("fatal") }, 0},
		{"unknown string == 5", func() int { return levelInt("unknown string") }, 5},
	}

	for _, t2 := range t2s {
		t.Run(t2.desc, func(t *testing.T) {
			if t2.f() != t2.want {
				t.Error("level code unmatch")
			}
		})
	}
}
