package process

import (
	"github.com/ctripcorp/nephele/command"

	"context"
	"errors"
	"fmt"
	"strings"
)

// 1. processString: "image/format,v_png/resize,w_100,h_200"
//
// 2. commandString: "image" | "format,v_png" | "resize,w_100,h_200"
//
// 3. word: "format" | "v_png" | "resize" | "w_100,h_200"
//
// 4. name: "format" | "resize"
//
// 5. k: "v" | "w" | "h"
//
// 6. v: "png" | "100" | "200"

func Parse(ctx context.Context, processString string) ([]command.Command, error) {
	commands := make([]command.Command, 0)
	if processString == "" {
		return commands, nil
	}
	commandStrings := strings.Split(processString, "/")
	for i, commandString := range commandStrings {
		if i == 0 {
			if commandString != "image" {
				return nil, errors.New("invalid process category: " + commandString)
			}
			continue
		}
		words := strings.Split(commandString, ",")
		name := words[0]
		cc, ok := command.List()[name]
		if !ok {
			return nil, fmt.Errorf("invalid command name: %s", name)
		}
		option := make(map[string]string)
		for j := 1; j < len(words); j++ {
			kv := strings.Split(words[j], "_")
			if len(kv) != 2 {
				continue
			}
			if kv[0] == "" {
				continue
			}
			if kv[1] == "" {
				continue
			}
			option[kv[0]] = kv[1]
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
	i := &image{
		state: imageIsBlob,
		blob:  blob,
		wand:  nil,
	}
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
