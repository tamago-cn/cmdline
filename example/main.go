package main

import (
	"os"

	"github.com/tamago-cn/cmdline"
)

func main() {
	cmdline.Interpret(os.Stdin)
}
