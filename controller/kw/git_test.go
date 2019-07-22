package kw

import "os"

func ExampleNewGitRepo() {
	os.Setenv("__KW_GIT_REPO", "/tmp/kwrepos")
	Sys("@rm -rfv /tmp/kwrepos")
	NewGitRepo("hello")
	Sys("ls /tmp/kwrepos")
	NewGitRepo("hello/world")
	Sys("ls /tmp/kwrepos")
	Sys("ls /tmp/kwrepos/hello")
	NewGitRepo("ignore/hello/all")
	Sys("ls /tmp/kwrepos/hello")
	// Output:
	//default
	// default
	// hello
	// world
	// all
	// world
}

func ExampleStore() {
	os.Setenv("__KW_GIT_REPO", "/tmp/kwrepos")
	os.Setenv("LC_ALL", "C")
	Sys("@rm -rfv /tmp/kwrepos")
	gr := NewGitRepo("hello/world")
	gr.Store("hello", []byte("world\n"))
	//log.SetLevel(log.TraceLevel)
	SysCd(gr.Dir, "ls")
	SysCd(gr.Dir, "cat hello")
	SysCd(gr.Dir, "git status")
	gr.Store("hello", []byte("updated world\n"))
	SysCd(gr.Dir, "ls")
	SysCd(gr.Dir, "cat hello")
	SysCd(gr.Dir, "sh -c", "git log --oneline | awk 'END { print NR}'")
	gr.Store("Dockerfile", []byte("FROM"))
	SysCd(gr.Dir, "ls")
	// Output:
	// hello
	// world
	// On branch master
	// nothing to commit, working tree clean
	// hello
	// updated world
	// 2
	// Dockerfile
	// hello

}
