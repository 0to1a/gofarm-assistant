package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	version              = "v1.1"
	versionNoProjectText = "GoFarm Assistant %s\n"
	versionText          = "GoFarm Assistant %s\nGoFarm project: %s\n"
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
	errFolderExist  = "Error: Folder exist"
	errFolderExist2 = "Error: Can't create inside GoFarm Folder"
	errFolderExist3 = "Error: Can't create module, no detect GoFarm Framework"
	okSuccess       = "Welcome to GoFarm family!"
	okSuccess2      = "Module %s created"
)

var (
	versionProject string
)

func checkProjectFile() bool {
	f, err := os.Open("main.go")
	if err != nil {
		fmt.Fprintln(output, err)
		return false
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "goFarmVersion =") {
			versionProject = strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), "goFarmVersion = ", ""), "\"", "")
			versionProject = strings.TrimSpace(versionProject)
			return true
		}
	}
	return false
}

func checkVersion() bool {
	files, err := filepath.Glob("*")
	if err != nil {
		fmt.Fprint(output, err)
		return false
	}

	versionProject = ""
	checklist := 0
	for _, filename := range files {
		if filename == "app" {
			checklist++
			continue
		}
		if filename == "framework" {
			checklist++
			continue
		}
		if filename == "main.go" {
			if checkProjectFile() {
				checklist++
			}
			continue
		}
	}

	if checklist == 3 && versionProject != "" {
		return true
	}
	return false
}

func printUsage() {
	fmt.Fprint(output, helpText)
}

func printVersion() {
	if checkVersion() {
		fmt.Fprintf(output, versionText, version, versionProject)
	} else {
		fmt.Fprintf(output, versionNoProjectText, version)
	}
}
