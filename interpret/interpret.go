package interpret

import (
	"github.com/gin-gonic/gin"
	"plugin"
	"strings"
)

var Config map[string]string

var interpreter func(string) (string, [][]string, error)

func Init() {
	if Config["type"] == "plugin" {
		p, err := plugin.Open(Config["path"])
		if err != nil {
			panic(err)
		}
		s, err := p.Lookup("Interpreter")
		if err != nil {
			panic(err)
		}
		interpreter = s.(func(string) (string, [][]string, error))
	}
}

func Do(c *gin.Context) (string, [][]string, error) {
	if interpreter != nil {
		return interpreter(c.Request.RequestURI)
	}
	key := c.Param("key")
	proccess := c.Query("x-nephele-process")
	return key, ParseProcess(proccess), nil
}

func ParseProcess(proc string) [][]string {
	var process [][]string
	cmds := strings.Split(proc, "/")
	for i, cmd := range cmds {
		if i == 0 {
			if cmd != "image" {
				return [][]string{}
			}
			continue
		}
		var command []string
		eles := strings.Split(cmd, ",")
		if eles[0] == "" {
			continue
		}
		command = append(command, eles[0])
		for j := 1; j < len(eles); j++ {
			kv := strings.Split(eles[j], "_")
			if len(kv) != 2 {
				continue
			}
			if kv[0] == "" {
				continue
			}
			if kv[1] == "" {
				continue
			}
			command = append(command, kv[0], kv[1])
		}
		process = append(process, command)
	}
	return process
}
