package kw

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// GitRepo manager
type GitRepo struct {
	repo *git.Repository
	dir  string
}

// NewGitRepo creates a new repo or open an existing one
// action shoud be a full path and all the parent directories must exist
func NewGitRepo(actionDir string) (gr *GitRepo, err error) {
	dir := filepath.Dir(actionDir)
	if _, e := os.Stat(dir); os.IsNotExist(e) {
		return nil, errors.New("not found parent directory " + dir)
	}
	// opening or creating repo
	var repo *git.Repository
	if _, e := os.Stat(actionDir); os.IsNotExist(e) {
		repo, err = git.PlainInit(actionDir, false)
	} else {
		repo, err = git.PlainOpen(actionDir)
	}
	if err == nil {
		return &GitRepo{repo, actionDir}, nil
	}
	return nil, err
}

// Store a file in the current git repo,
// creating it or updating if it does not exist
func (gr *GitRepo) Store(filename string, body []byte) (err error) {
	fullpath := filepath.Join(gr.dir, filename)
	if err = ioutil.WriteFile(fullpath, body, 0600); err != nil {
		return err
	}
	w, err := gr.repo.Worktree()
	if err != nil {
		return err
	}
	if _, err = w.Add(filename); err != nil {
		return err
	}
	_, err = w.Commit(filename, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Knative Whisk",
			Email: "knative@whisk",
			When:  time.Now(),
		},
	})
	return err

}
