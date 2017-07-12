package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"notabug.org/icock/goo/slice"
)

func compileSite(incremental bool) {
	mark()
	if incremental {
		log.Println("Incremental compilication is not implemented yet.")
		log.Println("Performing full compilication...")
	}
	clear()
	compilePages()
	compilePosts()
	removeEmptyDirectories()
	unmark()
}

var marker string = path.Join("_site", "_mg_dirty")

func mark() {
	var dirty *os.File
	dirty, err := os.Create(marker)
	defer dirty.Close()
	if err != nil {
		log.Println(err)
		log.Fatal("Failed to mark _site directory.")
	}
}

func unmark() {
	err := os.RemoveAll(marker)
	if err != nil {
		log.Println(err)
		log.Fatal("Failed to unmark _site directory.")
	}
}

func clear() {
	removeRecursivelyOrPanic(path.Join("_site", "index.html"))
	removeRecursivelyOrPanic(path.Join("_site", "tag"))

	var paths []string
	paths, err := filepath.Glob("_site/*/*/index.html") // _site/category/slug/index.html
	if err != nil {
		log.Println(err)
		log.Fatal("Failed to clear _site.")
	}
	slice.ForEachString(paths, removeRecursivelyOrPanic)
}

func removeRecursivelyOrPanic(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		log.Println(err)
		log.Fatal("Failed to remove " + path)
	}
}

func removeEmptyDirectories() {
	var categories []string = ls("_site")
	for _, category := range categories {
		var slugs []string = ls(filepath.Join("_site", category))
		for _, slug := range slugs {
			err := os.RemoveAll(filepath.Join("_site", category, slug))
			if err == nil {
				log.Print("Purging _site/" + category + "/" + slug)
			}
		}
	}
}

func ls(path string) []string {
	var file *os.File
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Println(err)
		log.Fatalln("Failed to open" + path + "directory.")
	}
	var directories []string
	directories, err =  file.Readdirnames(0)
	if err != nil {
		log.Println(err)
		log.Fatalln("Encounting errors when list contents of" + path + "directory.")
	}
	return directories
}