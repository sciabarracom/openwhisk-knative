package kw

import (
	"fmt"
	"os"
)

func ExampleNewGitRepo() {
	Sys("@sh -c", "rm -Rf /tmp/kwgit && mkdir /tmp/kwgit")
	_, err := NewGitRepo("/tmp/kwgit/hello")
	Sys("ls /tmp/kwgit")
	fmt.Println(err)
	_, err = NewGitRepo("/tmp/kwgit/world")
	Sys("ls /tmp/kwgit")
	fmt.Println(err)
	_, err = NewGitRepo("/tmp/kwgit/missing/hello")
	Sys("ls /tmp/kwgit")
	fmt.Println(err)
	// Output:
	// hello
	// <nil>
	// hello
	// world
	// <nil>
	// hello
	// world
	// not found parent directory /tmp/kwgit/missing
}

func Example_Store() {
	os.Setenv("LC_ALL", "C")
	Sys("@sh -c", "rm -rfv /tmp/kwgit && mkdir /tmp/kwgit")
	gr, err := NewGitRepo("/tmp/kwgit/hello")
	fmt.Println(err)
	fmt.Println(gr.Store("main", []byte("world\n")))
	//log.SetLevel(log.TraceLevel)
	SysCd(gr.dir, "ls")
	SysCd(gr.dir, "cat main")
	SysCd(gr.dir, "git status")
	gr.Store("main", []byte("updated world\n"))
	SysCd(gr.dir, "ls")
	SysCd(gr.dir, "cat main")
	SysCd(gr.dir, "sh -c", "git log --oneline | awk 'END { print NR}'")
	gr.Store("Dockerfile", []byte("FROM"))
	SysCd(gr.dir, "ls")
	// Output:
	// <nil>
	// <nil>
	// main
	// world
	// On branch master
	// nothing to commit, working tree clean
	// main
	// updated world
	// 2
	// Dockerfile
	// main
}

func Example_misc() {
	os.Setenv("LC_ALL", "C")
	Sys("@sh -c", "rm -rfv /tmp/kwgit")
}
