package core

import (
	"fmt"
	"os"

	"github.com/celestix/autodeployer_api/config"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func getRepoUrl(repoOwner, repoPath string) string {
	return fmt.Sprintf("https://github.com/%s/%s", repoOwner, repoPath)
}

const DEF_REMOTE = "origin"

func CloneRepository(projectName, repoOwner, repoName, branch, gho string) error {
	dir := fmt.Sprintf("%s/%s", config.Data.DataDirectory, projectName)
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		return fmt.Errorf("directory %s already exists", dir)
	}
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		RemoteName:        DEF_REMOTE,
		URL:               getRepoUrl(repoOwner, repoName),
		ReferenceName:     plumbing.NewBranchReferenceName(branch),
		SingleBranch:      true,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Auth: &http.BasicAuth{
			Username: "oauth2",
			Password: gho,
		},
	})
	return err
}

func PullRepository(projectName, gho string) error {
	dir := fmt.Sprintf("%s/%s", config.Data.DataDirectory, projectName)
	fmt.Println("pulling path: ", dir)
	r, err := git.PlainOpen(dir)
	if err != nil {
		fmt.Println("failed to plain open", err)
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		fmt.Println("failed to get worktree", err)
		return err
	}

	err = w.Pull(&git.PullOptions{
		RemoteName: DEF_REMOTE,
		Auth: &http.BasicAuth{
			Username: "oauth2",
			Password: gho,
		},
	})
	if err != git.NoErrAlreadyUpToDate {
		return err
	}
	return nil
}
