package main

import (
	"io"
	"os"
)

var output io.Writer = os.Stdout

func main() {
	if len(os.Args) == 1 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "help":
		printUsage()
	case "version":
		printVersion()
	default:
		printUsage()
	}
}
