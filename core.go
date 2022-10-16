package main

import (
	"embed"
	"fmt"
	git "github.com/go-git/go-git/v5"
	"os"
	"strings"
)

//go:embed skeleton/*.template
var fsSkeleton embed.FS

func checkCriteriaNewProject(folderPath string) {
	if _, err := os.Stat(folderPath); !os.IsNotExist(err) {
		fmt.Fprint(output, errFolderExist)
		os.Exit(3)
	}
	if checkVersion() {
		fmt.Fprint(output, errFolderExist2)
		os.Exit(3)
	}
}

func checkCriteriaNewModule(folderPath string) {
	if _, err := os.Stat(folderPath); !os.IsNotExist(err) {
		fmt.Fprint(output, errFolderExist)
		os.Exit(3)
	}
	if checkVersion() {
		fmt.Fprint(output, errFolderExist2)
		os.Exit(3)
	}
}

func removeUnusedNewProject(folderPath string) {
	os.RemoveAll(folderPath + "/.git")
	os.RemoveAll(folderPath + "/.github")
	os.RemoveAll(folderPath + "/.gitignore")
	os.RemoveAll(folderPath + "/CODE_OF_CONDUCT.md")
	os.RemoveAll(folderPath + "/LICENSE")
	os.RemoveAll(folderPath + "/README.md")
}

func createProject(folderPath string, projectName string) {
	if projectName == "" {
		fmt.Fprint(output, "Enter Project Name: ")
		fmt.Scanln(&projectName)
		folderPath = projectName
	}
	projectName = strings.TrimSpace(projectName)
	folderPath = strings.TrimSpace(folderPath)
	if projectName == "" || folderPath == "" {
		fmt.Fprint(output, errFolderExist)
		return
	}
	checkCriteriaNewProject(folderPath)

	_, err := git.PlainClone(folderPath, false, &git.CloneOptions{
		URL:      "https://github.com/0to1a/gofarm",
		Progress: output,
	})
	if err != nil {
		fmt.Fprint(output, err)
		return
	}

	removeUnusedNewProject(folderPath)

	read, err := os.ReadFile(folderPath + "/main.go")
	if err != nil {
		fmt.Fprint(output, err)
		return
	}
	newContents := strings.ReplaceAll(string(read), "ProjectName", projectName)
	err = os.WriteFile(folderPath+"/main.go", []byte(newContents), 0)
	if err != nil {
		fmt.Fprint(output, err)
		return
	}

	fmt.Fprint(output, okSuccess)
}

func createModule(moduleName string) {
	if moduleName == "" {
		fmt.Fprint(output, "Enter Module Name: ")
		fmt.Scanln(&moduleName)
	}
	moduleName = strings.ReplaceAll(moduleName, " ", "")
	if moduleName == "" {
		fmt.Fprint(output, errFolderExist)
		return
	}
	path := "app/" + moduleName + "Module"
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		fmt.Fprint(output, errFolderExist)
		return
	}
	if !checkVersion() {
		fmt.Fprint(output, errFolderExist3)
		return
	}

	d, err := fsSkeleton.ReadDir("skeleton")
	if err != nil {
		fmt.Fprint(output, err)
		return
	}

	os.Mkdir(path, os.ModePerm)
	path += "/"
	for _, entry := range d {
		fileData, err := fsSkeleton.ReadFile("skeleton/" + entry.Name())
		if err != nil {
			fmt.Fprint(output, err)
			return
		}
		newName := strings.ReplaceAll(entry.Name(), ".template", "")
		replacement := strings.ReplaceAll(string(fileData), "TEMPLATE", moduleName)
		f, _ := os.Create(path + newName)
		f.WriteString(replacement)
		f.Close()
		fmt.Fprintln(output, "copy: ", newName)
	}
	os.Mkdir(path+"migration", os.ModePerm)
	os.Create(path + "migration/.gitkeep")
	fmt.Fprintf(output, okSuccess2, moduleName)
}
