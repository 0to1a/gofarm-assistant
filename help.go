package main

import (
	"fmt"
	"os"
)

const (
	version = "v1.1"
)

func printUsage() {
	help := `GoFarm Assistant CLI

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
	fmt.Print(help)
	os.Exit(2)
}

func printVersion() {
	fmt.Printf("GoFarm Assistant %s\n", version)
	// TODO: check folder version
}
