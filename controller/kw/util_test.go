package kw

import (
	"errors"
	"fmt"

	middleware "github.com/go-openapi/runtime/middleware"
)

func ExampleLogIf() {
	LogIf(nil)
	err := fmt.Errorf("generated error")
	LogIf(err)
	// Output:
	// WARN: generated error
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
	// ERR: I am on panic
	// &{I am on panic}
	// no panic
	// <nil>
	// ERR: Panic mode
	// Panic mode
}
