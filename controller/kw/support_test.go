package kw

import (
	"errors"
	"fmt"

	middleware "github.com/go-openapi/runtime/middleware"
)

func shouldPanicRest(flag bool) (resp middleware.Responder) {
	defer RecoverRest(&resp)
	if flag {
		err := errors.New("I am on panic")
		FatalIf(err)
	}
	return middleware.NotImplemented("not panicing")
}

func Example2FatalIf() {
	show(shouldPanicRest(false))
	show(shouldPanicRest(true))
	// Output:
	// &{501 not panicing map[]}
	// &{I am on panic}
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
