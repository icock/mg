package main

import (
	"fmt"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func gitRepository() *git.Repository {
	repo, err := git.PlainOpen(siteRoot)
	if err != nil {
		panic(fmt.Sprintf("Failed to open git repository at %s", siteRoot))
	}
	return repo
}

func gitLastCommitTime() time.Time {
	var head *plumbing.Reference
	head, err := siteRepository.Head()
	if err != nil {
		panic("Failed to get head of git repository!")
	}
	var hash plumbing.Hash = head.Hash()
	var commit *object.Commit
	commit, err = siteRepository.CommitObject(hash)
	if err != nil {
		panic(err)
	}
	return commit.Committer.When
}
