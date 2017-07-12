package main

import (
	"time"
	"os"
	"path/filepath"
	"notabug.org/icock/goo/slice"
	"fmt"
	"io/ioutil"
)


func layoutSkeleton(basePath string) {
	file, err := layoutFile(basePath)
	defer file.Close()
	var layoutTemplates = map[string][]byte{
		"header": []byte(`
<!DOCTYPE html>
<html lang="en">
<head>
	<title>{{title}}</title>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="stylesheet" href="https://icock.github.io/markgone/css/godoc.css">
</head>
<body>
<h1><a href="/">{{siteName}}</a></h2>
<main>
<article>
		`),
		"footer": []byte(`
</article>
</main>
<footer>
<nav>
<section class="page-list">
{{pages}}
</section>
<section class="category-list">
{{categories}}
</section>
<section class="tag-list">
{{tags}}
</section>
</nav>
<p>To report a problem with the web site, open an issue or send a pull request
at <a herf="{{repoUrl}}">source repository</a>.</p>
<p>Last Modified: <time>{{commitTime}}</time><p>
<p>The contents of this site, unless otherwise expressly stated,
are dedicated to the Public Domain.</p>
</footer>
</body>
</html>
		`),
	}

	if err == nil {
		err = ioutil.WriteFile(layoutPath(basePath), layoutTemplates[basePath], 0644)
		if err != nil {
			panic(fmt.Sprintf("Cannot write `_layout/%s.html` file", basePath))
		}
	}
}

func isLayoutModified(siteLastModified time.Time) bool {
	return slice.SomeString(
		[]string{"header", "footer"},
		func (basePath string) bool {
			return layoutLastModified(basePath).After(siteLastModified)
		})
}

func layoutLastModified(basePath string) time.Time {
	file, err := layoutFile(basePath)
	if err != nil {
		panic("Cannot read _layout/" + basePath)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		panic("Cannot read file info: _layout/" + basePath)
	}

	return fileInfo.ModTime()
}

func layoutFile(basePath string) (*os.File, error) {
	var path string = layoutPath(basePath)
	return os.Open(path)
}
func layoutPath(basePath string) string {
	return filepath.Join(siteRoot, "_layout", basePath+".html")
}
