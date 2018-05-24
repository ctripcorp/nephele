package process

import (
    "github.com/ctripcorp/nephele/command"

    _  "github.com/ctripcorp/nephele/command/format"

    "context"
    "fmt"
)

func Parse(ctx context.Context, process [][]string) ([]command.Command, error) {
    commands := make([]command.Command, 0)
	for _, cmd := range process {
		name := cmd[0]
		option := make(map[string]string)
		for i := 1; i < len(cmd); i = i + 2 {
			option[cmd[i]] = cmd[i+1]
		}
        cc, ok := command.List()[name]
        if !ok {
            return nil, fmt.Errorf("invalid command name: %s", name)
		}
        c := cc()
        err := c.Verify(ctx, option)
        if err != nil {
            return nil, err
        }
        commands = append(commands, c)
    }
    return commands, nil
}

func Do(ctx context.Context, blob []byte, commands []command.Command) ([]byte, error) {
	i := &image{imageIsBlob, blob, nil}
	for _, command := range commands {
        if command.Support() == "wand" {
	        w, err := i.Wand()
	        if err != nil {
                return nil, err
	        }
            err = command.ExecuteOnWand(ctx, w)
            if err != nil {
                return nil, err
            }
            continue
        }
        if command.Support() == "blob" {
            b, err := i.Bytes()
			if err != nil {
                return nil, err
			}
            b, err = command.ExecuteOnBlob(ctx, b)
            if err != nil {
                return nil, err
            }
            i.SetBlob(b)
            continue
        }
	}
	return i.Bytes()
}
