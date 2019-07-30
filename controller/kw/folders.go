package kw

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// FolderManager manage folders
type FolderManager struct {
	baseDir   string
	Namespace string
}

// NewFolderManager creates a folder manager specifing the default namespace
// KW_REPO environment variables points to the base repo
// defaults to "/tmp/kw/repo"
func NewFolderManager(defaultNamespace string) *FolderManager {
	baseDir := os.Getenv("KW_REPO")
	if baseDir == "" {
		baseDir = "/tmp/kw/repo"
	}
	if !ValidateURLPathComponent(defaultNamespace) {
		panic("invalid namespace")
	}
	// ensure we have baseDir, default namespace and default package
	path := filepath.Join(baseDir, defaultNamespace, "default")
	if err := os.MkdirAll(path, 0755); err != nil {
		panic("cannot create dir " + path)
	}
	return &FolderManager{baseDir, defaultNamespace}
}

// ListNamespaces returns the current namespace
func (fm *FolderManager) ListNamespaces() []string {
	return []string{fm.Namespace}
}

// ListPackages return the list of the current packages
func (fm *FolderManager) ListPackages() []string {
	dir := filepath.Join(fm.baseDir, fm.Namespace)
	infos, err := ioutil.ReadDir(dir)
	PanicIf(err)
	res := []string{}
	for _, info := range infos {
		if info.IsDir() && info.Name() != "default" {
			res = append(res, info.Name())
		}
	}
	sort.Strings(res)
	return res
}

// UpdatePackage create a folder for packages if it does not exists;
// you cannot use "default" as a package name
func (fm *FolderManager) UpdatePackage(name string) error {
	if name == "default" {
		return errors.New("unable to create package 'default': package name 'default' is reserved")
	}
	if !ValidateURLPathComponent(name) {
		return fmt.Errorf("unable to create package '%s': the name of the entity contains illegal characters", name)
	}
	dir := filepath.Join(fm.baseDir, fm.Namespace, name)
	return os.MkdirAll(dir, 0755)
}

// ListActions return a list of actions in a given package
func (fm *FolderManager) ListActions(packge *string) []string {
	var packges []string
	if packge == nil {
		packges = fm.ListPackages()
		packges = append(packges, "default")
	} else {
		packges = []string{*packge}
	}
	result := []string{}
	for _, packge := range packges {
		dir := filepath.Join(fm.baseDir, fm.Namespace, packge)
		infos, err := ioutil.ReadDir(dir)
		LogIf(err)
		if err != nil {
			continue
		}
		for _, info := range infos {
			if info.IsDir() {
				var name string
				if packge == "default" {
					name = fmt.Sprintf("/%s/%s", fm.Namespace, info.Name())
				} else {
					name = fmt.Sprintf("/%s/%s/%s", fm.Namespace, packge, info.Name())
				}
				result = append(result, name)
			}
		}
	}
	sort.Strings(result)
	return result
}

// splitActionName splits action name in components
func (fm *FolderManager) splitActionName(actionName string) (namespace string, packge string, action string, err error) {
	namespace = fm.Namespace
	packge = "default"
	err = nil
	a := strings.Split(actionName, "/")
	switch n := len(a); n {
	case 1:
		action = a[0]
	case 2:
		packge = a[0]
		action = a[1]
	default:
		namespace = a[n-3]
		packge = a[n-2]
		action = a[n-1]
		if (n == 4 && a[0] != "") || n > 4 {
			err = errors.New("the requested resource was not found")
		}
	}
	if namespace == "_" {
		namespace = fm.Namespace
	}
	//log.Debug(namespace, packge, action)
	if !ValidateURLPathComponent(namespace) {
		err = fmt.Errorf("namespace '%s' contains illegal characters", namespace)
	}
	if !ValidateURLPathComponent(packge) {
		err = fmt.Errorf("package '%s' contains illegal characters", packge)
	}
	if !ValidateURLPathComponent(action) {
		err = fmt.Errorf("action '%s' contains illegal characters", action)
	}
	return
}

// UpdateAction create a git repo for an action
// the action name can optionallu include the package name that defaults to "default"
// and the namespace that defaults to the default namespace
func (fm *FolderManager) UpdateAction(name string) (*GitRepo, error) {
	namespace, packge, action, err := fm.splitActionName(name)
	if err != nil {
		return nil, err
	}
	pkgDir := filepath.Join(fm.baseDir, namespace, packge)
	if _, err := os.Stat(pkgDir); err != nil {
		return nil, fmt.Errorf("packge %s in namespace %s not found", packge, namespace)
	}
	actionDir := filepath.Join(pkgDir, action)
	return NewGitRepo(actionDir)
}
