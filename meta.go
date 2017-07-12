package main

import (
	"os"
	"path/filepath"
	"strings"
	"notabug.org/icock/goo/slice"
)

// panics outside a git repository
func gitRoot() string {
	var wd string
	wd, _ = os.Getwd()
	if isDirectoryAvailable(".git") {
		return wd
	} else {
		if isSystemRoot(wd) {
			panic(newFileSystemNotFoundException(
				"`mg` should be run within a git repository."))
		} else {
			os.Chdir("..")
			return gitRoot()
		}
	}
}
func isSystemRoot(rootedPath string) bool {
	var path string = filepath.ToSlash(rootedPath)
	if strings.HasSuffix(path, "/") { // "/" on unix, "C:/" and etc. on windows
		return true
	} else {
		return false
	}
}

// Ignore `tag/` and directories that starts with `.` or `_`.
func getCategories(gitRoot *os.File) []string {
	var directories []string
	var err error
	directories, err = gitRoot.Readdirnames(0)
	if err != nil {
		panic(newFileSystemException("Cannot read git root directory."))
	}
	return slice.FilterString(directories, func (d string) bool {
		return !strings.HasPrefix(d, ".") && !strings.HasPrefix(d, "_") && d != "tag"
	})
}