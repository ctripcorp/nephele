package main

import (
	"fmt"
	"github.com/ctripcorp/nephele/command"
	_ "github.com/ctripcorp/nephele/command/format"
)

func main() {
	fmt.Println(command.List())
}
