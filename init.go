package main

import (
	"os"
	"path/filepath"
	"fmt"
)

func initSite(categories []string) {
	initLayout()

	err := os.MkdirAll("_site", 0700)
	if err != nil {
		panic(newFileSystemException("Cannot create output directory `_site`."))
	}

	for _, category := range categories {
		categoryPath := filepath.Join(siteRoot, "_site", category)
		os.MkdirAll(categoryPath, 0700)
	}
}

func initLayout() {
	err := os.MkdirAll("_layout", 0700)
	if err != nil {
		panic(newFileSystemException("Cannot create directory `_layout`."))
	}

	layoutSkeleton("header")
	layoutSkeleton("footer")
}



