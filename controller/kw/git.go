package kw

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// GitRepo manager
type GitRepo struct {
	Package string
	Action  string
	Repo    *git.Repository
	Dir     string
}

func baseDir() string {
	// initialize basedir
	base := os.Getenv("__KW_GIT_REPO")
	if base == "" {
		base = "/var/lib/kw/repo"
	}
	os.MkdirAll(base, 0755)
	return base
}

// NewGitRepo creates a new repo or open an existing one
// action shoud be in format "action" or "package/action"
func NewGitRepo(action string) (gr *GitRepo) {

	// paths to get git repo
	base := baseDir()
	paths := strings.Split(action, "/")
	n := len(paths)
	if n > 2 {
		log.Warn("too many slashes, ignoring leading path components")
	}

	act := paths[n-1]
	pkg := "default"
	if n > 1 {
		pkg = paths[n-2]
	}

	// making directory
	pkgDir := filepath.Join(base, pkg)
	actDir := filepath.Join(pkgDir, act)
	if _, err := os.Stat(pkgDir); os.IsNotExist(err) {
		os.MkdirAll(pkgDir, 0700)
	}

	// opening or creating repo
	var repo *git.Repository
	var err error
	if _, e := os.Stat(actDir); os.IsNotExist(e) {
		repo, err = git.PlainInit(actDir, false)
	} else {
		repo, err = git.PlainOpen(actDir)
	}
	PanicIf(err)
	return &GitRepo{pkg, action, repo, actDir}
}

// Store a file in the current git repo,
// creating it or updating if it does not exist
func (gr *GitRepo) Store(filename string, body []byte) {
	var err error
	fullpath := filepath.Join(gr.Dir, filename)
	err = ioutil.WriteFile(fullpath, body, 0600)
	PanicIf(err)
	w, err := gr.Repo.Worktree()
	PanicIf(err)
	_, err = w.Add(filename)
	PanicIf(err)
	_, err = w.Commit(filename, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Knative Whisk",
			Email: "knative@whisk",
			When:  time.Now(),
		},
	})
	PanicIf(err)
}
