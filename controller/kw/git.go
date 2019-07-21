package kw

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
)

// GitRepo manager
type GitRepo struct {
	pkg    string
	action string
	repo   *git.Repository
}

func baseDir() string {
	// initialize basedir
	base := os.Getenv("KW_REPO_BASE")
	if base == "" {
		base = "/tmp/kwrepos"
	}
	os.MkdirAll(base, 0755)
	return base
}

// NewGitRepo pcreates a new handler
// providing the base where git repos are stored
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

	pkgDir := base + "/" + pkg
	actDir := pkgDir + "/" + act

	// making directory
	os.MkdirAll(pkgDir, 0700)
	repo, err := git.PlainInit(actDir, false)
	PanicIf(err)
	return &GitRepo{pkg, action, repo}
}
