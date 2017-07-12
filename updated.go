package main

import (
	"time"
	"path/filepath"
	"os"
	"bufio"
	"strings"
	"fmt"
)


func needFullUpdate(siteLastModified time.Time) bool {
	if siteLastModified.IsZero() { // first time
		return true
	} else if isLayoutModified(siteLastModified) {
		return true
	} else {
		return false
	}
}


func isSiteUpdated(siteLastModified time.Time) bool {
	var gitLastCommit time.Time = gitLastCommitTime()
	if siteLastModified.Equal(gitLastCommit) {
		return true
	} else if siteLastModified.After(gitLastCommit) {
		panic(
			"siteLastModified > gitLastCommit:\n" +
				"there may be issues on the git repository or system time")
	} else {
		return false
	}
}

// Returns site last modified time via checking _site/humans.txt.
// Returns zero time (0001-01-01T00:00:00Z) if humans.txt does not exist.
// Panics if humans.txt dose not contain a line starting with `Last Modified: `.
func siteLastModifiedTime() time.Time {
	humansTxtPath := filepath.Join(siteRoot, "_site", "humans.txt")
	var err error
	var humansTxt *os.File
	humansTxt, err = os.Open(humansTxtPath)
	defer humansTxt.Close()
	if err != nil {
		return time.Time{}
	}

	lastModifiedPrefix := "Last Modified: "
	rfc3339 := "2006-06-30T23:59:60Z"
	scanner := bufio.NewScanner(humansTxt)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, lastModifiedPrefix) {
			line = strings.TrimPrefix(line, lastModifiedPrefix)
			var date time.Time
			date, err := time.Parse(rfc3339, line)
			if err != nil {
				panic(fmt.Sprintf("Failed to parse time of in humans.txt\n%s", line))
			}
			return date
		}
	}
	panic("humans.txt does not contain last modified time.")
}

