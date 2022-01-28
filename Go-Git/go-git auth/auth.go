package main

import (
	"fmt"
	_ "log"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

func main() {
	cloneGit()
}
func cloneGit() (err error) {
	username := "Vasanth-rasf"
	password := "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"
	repo := "github.com/Vasanth-rasf/go-git-saplre.git"
	url := fmt.Sprintf("https://%s:%s@%s", username, password, repo)
	options := git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	}
	r, err := git.PlainClone("./go-git_auth", false, &options)

	if err != nil {
		return err
	}
	fmt.Println(r.Log)
	return err
}
