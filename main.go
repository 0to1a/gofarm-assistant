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
	case "new":
		if len(os.Args) == 4 {
			createProject(os.Args[2], os.Args[3])
		} else if len(os.Args) == 3 {
			createProject(os.Args[2], os.Args[2])
		} else {
			createProject("", "")
		}
	case "module":
		if len(os.Args) == 3 {
			createModule(os.Args[2])
		} else {
			createModule("")
		}
	case "help":
		printUsage()
	case "version":
		printVersion()
	default:
		printUsage()
	}
}
