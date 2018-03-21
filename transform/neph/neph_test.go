package neph

import "testing"
import "github.com/ctripcorp/nephele/transform/command"

func TestAccept(t *testing.T) {
	type Check struct {
		Name  string
		Param map[string]string
		IsErr bool
	}
	checks := []Check{
		Check{Name: command.RESIZE, Param: map[string]string{"w": "100", "h": "100"}, IsErr: false},
		Check{Name: command.RESIZE, Param: map[string]string{"w": "abc", "h": "100"}, IsErr: true},
		Check{Name: command.RESIZE, Param: map[string]string{"m": "fixed"}, IsErr: true},
	}
	transformer := &Transformer{}
	for i, c := range checks {
		err := transformer.Accept(c.Name, c.Param)
		if c.IsErr != (err != nil) {
			t.Error("index:", i, ",Name:", c.Name, " check failed!")
		}
	}
}
