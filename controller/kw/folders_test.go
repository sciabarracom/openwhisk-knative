package kw

func ExampleNewFolderManager_panic() {
	SysSh("@rm -Rf /tmp/kw ; touch /tmp/kw")
	capture(func() {
		NewFolderManager("invalid!")
	})
	capture(func() {
		NewFolderManager("kwhisk")
	})
	// Output:
	// capture: invalid namespace
	// capture: cannot create dir /tmp/kw/repo/kwhisk/default
}

func ExampleSplitActionName() {
	SysSh("@rm -Rf /tmp/kw")
	fm := NewFolderManager("kwhisk")
	print(fm.splitActionName("hello"))
	print(fm.splitActionName("hello/world"))
	print(fm.splitActionName("michele/hello/world"))
	print(fm.splitActionName("/michele/hello/world"))
	print(fm.splitActionName("_/hello/world"))
	print(fm.splitActionName("/_/hello/world"))
	print(fm.splitActionName("etc/michele/hello/world"))
	print(fm.splitActionName("/etc/michele/hello/world"))
	print(fm.splitActionName("//michele/hello/world"))
	print(fm.splitActionName("hello!"))
	print(fm.splitActionName("hello!/world"))
	print(fm.splitActionName("michele!/hello/world"))
	// Output:
	// kwhisk default hello <nil>
	// kwhisk hello world <nil>
	// michele hello world <nil>
	// michele hello world <nil>
	// kwhisk hello world <nil>
	// kwhisk hello world <nil>
	// michele hello world the requested resource was not found
	// michele hello world the requested resource was not found
	// michele hello world the requested resource was not found
	// kwhisk default hello! action 'hello!' contains illegal characters
	// kwhisk hello! world package 'hello!' contains illegal characters
	// michele! hello world namespace 'michele!' contains illegal characters
}

func ExampleFolderManager() {
	SysSh("@rm -Rf /tmp/kw")
	fm := NewFolderManager("kwhisk")
	print(fm.ListNamespaces())
	print(fm.ListPackages())
	fm.UpdatePackage("hello")
	print(fm.ListPackages())
	fm.UpdatePackage("world")
	print(fm.ListPackages())
	print(fm.ListActions(nil))
	s := "donotexist"
	print(fm.ListActions(&s))
	gr, err := fm.UpdateAction("hello")
	print(err, gr.dir)
	print(fm.ListActions(nil))
	gr, err = fm.UpdateAction("hello/world")
	print(err, gr.dir)
	print(fm.ListActions(nil))
	fm.UpdateAction("hello/hello")
	fm.UpdatePackage("world")
	fm.UpdateAction("world/world")
	print(fm.ListPackages())
	print(fm.ListActions(nil))
	// Output:
	// [kwhisk]
	// []
	// [hello]
	// [hello world]
	// []
	// []
	// <nil> /tmp/kw/repo/kwhisk/default/hello
	// [/kwhisk/hello]
	// <nil> /tmp/kw/repo/kwhisk/hello/world
	// [/kwhisk/hello /kwhisk/hello/world]
	// [hello world]
	// [/kwhisk/hello /kwhisk/hello/hello /kwhisk/hello/world /kwhisk/world/world]
}
