package main

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

// Example of how to:
// - Clone a repository into memory
// - Get the HEAD reference
// - Using the HEAD reference, obtain the commit this reference is pointing to
// - Using the commit, obtain its history and print it
func main() {

	cmtid := "a1dcd5b5f28c5b13468497daf7a4ce4122406207"
	// Clones the given repository, creating the remote, the local branches
	// and fetching the objects, everything in memory:
	// Clones the given repository in memory, creating the remote, the local
	// branches and fetching the objects, exactly as:
	Info("git clone https://github.com/Vasanth-rasf/gitlearning.git ")

	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/Vasanth-rasf/gitlearning.git",
	})

	CheckIfError(err)

	// Gets the HEAD history from HEAD, just like this command:
	Info("git log")

	// ... retrieves the branch pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)

	// ... retrieves the commit history
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	CheckIfError(err)

	// ... just iterates over the commits, printing it
	err = cIter.ForEach(func(c *object.Commit) error {
		//fmt.Println(c.Hash)
		if c.Hash.String() == cmtid {
			fmt.Println("Your commit id is there")

		}
		return nil
	})
	CheckIfError(err)
}
