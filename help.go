package main

import (
	"fmt"
)

const (
	version              = "v1.1"
	versionNoProjectText = "GoFarm Assistant %s\n"
	versionText          = "GoFarm Assistant %s\n  GoFarm project: %s\n"
	helpText             = `GoFarm Assistant CLI

Usage: 
  gofarm-assistant [command] [arg...] 

Commands:
  new		Create new project
  module	Generate new module app on your project folder
  fix		Check & auto-repair before build application
  generate	Scan from database and generate struct
  upgrade	Upgrade project to latest framework
  help		Help about any command
  version	Print current GoFarm Assistant version & Project version
`
)

func printUsage() {
	fmt.Fprint(output, helpText)
}

func printVersion() {
	fmt.Fprintf(output, versionNoProjectText, version)
	// TODO: check folder version
}
