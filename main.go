package main

import "flag"

func main() {
	if len(flag.Args()) < 1 {
		printUsage()
	}

	switch flag.Arg(0) {
	case "create":
		//printVersion()
	case "help":
		printUsage()
	case "version":
		printVersion()
	}
}
