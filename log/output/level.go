package output

var levels []string = []string{
	"fatal",
	"error",
	"warning",
	"info",
	"debug",
}

func levelInt(level string) int {
	var i = 0
	for ; i < 5; i++ {
		if levels[i] == level {
			return i
		}
	}
	return 3
}
