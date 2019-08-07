package kw

import (
	"errors"
	"fmt"
)

func ExampleLogIf() {
	fmt.Println(LogIf(nil))
	err := fmt.Errorf("generated error")
	if LogIf(err) {
		fmt.Println(LastError)
	}
	// Output:
	// false
	// generated error
}

func shouldPanic(flag bool) (err error) {
	defer Recover(&err)
	if flag {
		err := errors.New("Panic mode")
		FatalIf(err)
	}
	fmt.Println("no panic")
	return nil
}

func ExampleFatalIf() {
	fmt.Println(shouldPanic(false))
	fmt.Println(shouldPanic(true))
	// Output:
	// no panic
	// <nil>
	// Panic mode
}

func ExampleSys() {
	Sys("/bin/echo 1 2 3")
	Sys("/bin/echo 3", "4", "5")
	Sys("@sh -c", "echo foo >/tmp/foo")
	fmt.Print(Sys("cat /tmp/foo"))
	Sys("@sh -c", "echo bar >/tmp/bar")
	fmt.Print(Sys("@cat /tmp/bar"))
	// Output:
	// 1 2 3
	// 3 4 5
	// foo
	// foo
	// bar
}

func ExampleSysCd() {
	Sys("@mkdir -p /tmp/example-syscd/hello-i-am-a-dir")
	fmt.Println(SysCd("/tmp/example-syscd", "@ls"))
	// Output:
	// hello-i-am-a-dir
}

func ExampleSysSh() {
	show(SysSh("@rm -Rf /tmp/simple ; mkdir /tmp/simple ; ls -d /tmp/simple"))
	SysSh("cd /tmp/simple ; touch hello ; ls")
	// Output:
	// /tmp/simple
	//
	// hello
}
