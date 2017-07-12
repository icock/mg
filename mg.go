package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"gopkg.in/src-d/go-git.v4"
)

var siteRoot string = gitRoot()
var siteRepository *git.Repository = gitRepository()

func main() {
	var siteRootDirectory *os.File
	siteRootDirectory, err := os.Open(siteRoot)
	defer siteRootDirectory.Close()
	if err != nil {
		panic(newFileSystemException(fmt.Sprintf("Cannot create open %s", siteRoot)))
	}
	var categories []string = getCategories(siteRootDirectory)

	os.Chdir(siteRoot)
	initSite(categories)

	var siteLastModified time.Time = siteLastModifiedTime()
	if needFullUpdate(siteLastModified) {
		compileSite(false)
		deploySite()
	} else {
		if isSiteUpdated(siteLastModified) {
			log.Println("Everything is updated.")
		} else {
			compileSite(true)
			deploySite()
		}
	}
}




