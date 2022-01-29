package main

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func main() {

	r, err := git.PlainClone("./Branchoutput", false, &git.CloneOptions{
		URL:           "https://github.com/Vasanth-rasf/gitlearning.git",
		ReferenceName: plumbing.ReferenceName("refs/heads/develop"),
		SingleBranch:  true,
	})

	if err != nil {
		log.Fatal(err)
	}

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()

	if err != nil {
		log.Fatal(err)
	}

	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(commit)
}
