package main

import (
	"os"

	"github.com/shifty21/scone/command"
)

func main() {
	os.Exit(command.Run(os.Args[1:]))
}
