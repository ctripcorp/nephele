package neph

import (
	"strings"

	"github.com/ctripcorp/nephele/context"
	"github.com/ctripcorp/nephele/index"
	nephIndex "github.com/ctripcorp/nephele/index/neph"
	"github.com/ctripcorp/nephele/transform"
	nephTransform "github.com/ctripcorp/nephele/transform/neph"
)

//Decoder represents decoder
type Decoder struct {
	ctx         *context.Context
	transformer transform.Transformer
	path        string
}

//Decode uri
func (e *Decoder) Decode(uri string) error {
	e.path = strings.Split(uri, "?")[0]
	e.transformer = &nephTransform.Transformer{}
	processes := decode(uri)
	for _, process := range processes {
		if err := e.transformer.Accept(e.ctx, process.Name, process.Param); err != nil {
			return err
		}
	}
	return nil
}

type process struct {
	Name  string
	Param map[string]string
}

func decode(uri string) []process {
	procs := strings.Split(uri, "?")
	prefix := "x-nephele-process=image/"
	processes := []process{}
	for _, proc := range procs {
		if strings.HasPrefix(proc, prefix) {
			commands := strings.Split(strings.TrimPrefix(proc, prefix), "/")
			for _, command := range commands {
				if command == "" {
					continue
				}
				arr := strings.Split(command, ",")
				paramMap := make(map[string]string)
				for index := 1; index < len(arr); index++ {
					kv := strings.Split(arr[index], "_")
					if len(kv) != 2 {
						continue
					}
					paramMap[kv[0]] = kv[1]
				}
				processes = append(processes, process{Name: arr[0], Param: paramMap})
			}
		}
	}
	return processes
}

//CreateIndex create index
func (e *Decoder) CreateIndex() index.Index {
	return &nephIndex.Index{Ctx: e.ctx, Path: e.path}
}

//Transformer  transformer
func (e *Decoder) Transformer() transform.Transformer {
	return e.transformer
}
