package main

import (
	"os"
	"strings"
	"path/filepath"
	"log"
	"notabug.org/icock/goo/slice"
)

func deploySite() {
	err := os.Chdir("_site")
	if err != nil {
		panic("Cannot change into _site to deploy site!")
	}
	defer func() {
		err := os.Chdir(siteRoot)
		if err != nil {
			panic("Cannot change back to " + siteRoot + " directory!")
		}
	}()

	deployExecutables := []string{
		"deploy", // unix
		"deploy.exe", // windows
	}
	i := slice.IndexString(deployExecutables, isFileAvailable)
	if i != -1 {
		command := filepath.Join(siteRoot, "_site", deployExecutables[i])
		runDeployScript(command, []string{})
	} else if slice.SomeString([]string{"makefile", "Makefile", "GNUmakefile"}, isFileAvailable) {
		runDeployScript("make", []string{})
	} else if isFileAvailable("make.go") {
		runDeployScript("go", []string{"run", "make.go"})
	} else if isDirectoryAvailable(".git") {
		runDeployScript("git", []string{"push"})
	} else {
		log.Fatal("Do not know how to deploy _site.\nPlease deploy it manually.")
	}
}

func isFileAvailable(name string) bool {
	fileInfo, err := os.Stat(name)
	if err == nil {
		if fileInfo.IsDir() {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}
func isDirectoryAvailable(name string) bool {
	fileInfo, err := os.Stat(name)
	if err == nil {
		if fileInfo.IsDir() {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func runDeployScript(command string, args []string) {
	err := execCommand(command, args)
	if err != nil {
		panic("Deploy _site via " + command + strings.Join(args, " ") + " failed!")
	}
}
