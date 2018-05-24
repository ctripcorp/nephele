package interpret

import "testing"
import "fmt"

func TestParseProcess(t *testing.T) {

	var compare = func(a [][]string, b [][]string) bool {
		return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
	}

	var cases []string = []string{
		"image/resize,w_200,h_100",
		"image/resize,w_200,h_100/rotate,v_90",
		"image/,_/rotate,v_90",
	}

	var results [][][]string = [][][]string{
		[][]string{
			[]string{"resize", "w", "200", "h", "100"},
		},
		[][]string{
			[]string{"resize", "w", "200", "h", "100"},
			[]string{"rotate", "v", "90"},
		},
		[][]string{
			[]string{"rotate", "v", "90"},
		},
	}

	for i, c := range cases {
		if !compare(ParseProcess(c), results[i]) {
			t.Errorf("case: %s not successfully parsed", c)
		}
	}
}
