package kw

import (
	"errors"
	"fmt"

	middleware "github.com/go-openapi/runtime/middleware"
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

func shouldPanicRest(flag bool) (resp middleware.Responder) {
	defer RecoverRest(&resp)
	if flag {
		err := errors.New("I am on panic")
		PanicIf(err)
	}
	return middleware.NotImplemented("not panicing")
}

func shouldPanic(flag bool) (err error) {
	defer Recover(&err)
	if flag {
		err := errors.New("Panic mode")
		PanicIf(err)
	}
	fmt.Println("no panic")
	return nil
}

func ExamplePanicIf() {
	fmt.Println(shouldPanicRest(false))
	fmt.Println(shouldPanicRest(true))
	fmt.Println(shouldPanic(false))
	fmt.Println(shouldPanic(true))
	// Output:
	// &{501 not panicing map[]}
	// &{I am on panic}
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
	print(SysSh("@rm -Rf /tmp/simple ; mkdir /tmp/simple ; ls -d /tmp/simple"))
	SysSh("cd /tmp/simple ; touch hello ; ls")
	// Output:
	// /tmp/simple
	//
	// hello
}

func ExampleValidateURLPathComponent() {
	fmt.Println(ValidateURLPathComponent("th1s"))
	fmt.Println(ValidateURLPathComponent("th1s_is-"))
	fmt.Println(ValidateURLPathComponent("th1s_is-a.V4lid"))
	fmt.Println(ValidateURLPathComponent("th1s_is-a.V4lid~component"))
	fmt.Println(ValidateURLPathComponent(""))
	fmt.Println(ValidateURLPathComponent("not valid"))
	fmt.Println(ValidateURLPathComponent("not/valid"))
	fmt.Println(ValidateURLPathComponent("not%valid"))
	// Output:
	// true
	// true
	// true
	// true
	// false
	// false
	// false
	// false
}

func ExampleMkErr() {
	sp.Dump(MkErr(errors.New("wrong")))
	// Output:
	// (*models.ErrorMessage)({
	//  Code: (string) (len=32) "2bda2998d9b0ee197da142a0447f6725",
	//  Error: (*string)((len=5) "wrong")
	// })
}
