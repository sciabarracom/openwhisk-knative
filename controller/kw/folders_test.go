package kw

func ExampleNewFolderManager_panic() {
	SysSh("@rm -Rf /tmp/kwtest ; touch /tmp/kwtest")
	capture(func() {
		NewFolderManager("invalid!")
	})
	capture(func() {
		NewFolderManager("kwhisk")
	})
	// Output:
	// capture: invalid namespace
	// capture: cannot create dir /tmp/kwtest/kwhisk/default
}

func ExampleSplitActionName() {
	SysSh("@rm -Rf /tmp/kwtest")
	fm := NewFolderManager("kwhisk")
	show(fm.SplitActionName("hello"))
	show(fm.SplitActionName("hello/world"))
	show(fm.SplitActionName("michele/hello/world"))
	show(fm.SplitActionName("/michele/hello/world"))
	show(fm.SplitActionName("_/hello/world"))
	show("---")
	show(fm.SplitActionName("/hello"))
	show(fm.SplitActionName("/_/hello"))
	show(fm.SplitActionName("/_/hello/world"))
	show("---")
	show(fm.SplitActionName("etc/michele/hello/world"))
	show(fm.SplitActionName("/etc/michele/hello/world"))
	show(fm.SplitActionName("//michele/hello/world"))
	show("---")
	show(fm.SplitActionName("hello!"))
	show(fm.SplitActionName("hello!/world"))
	show(fm.SplitActionName("michele!/hello/world"))
	// Output:
	// kwhisk default hello <nil>
	// kwhisk hello world <nil>
	// michele hello world <nil>
	// michele hello world <nil>
	// kwhisk hello world <nil>
	// ---
	// kwhisk default hello <nil>
	// kwhisk default hello <nil>
	// kwhisk hello world <nil>
	// ---
	// kwhisk default  action '' contains illegal characters
	// kwhisk default  action '' contains illegal characters
	// kwhisk default  action '' contains illegal characters
	// ---
	// kwhisk default hello! action 'hello!' contains illegal characters
	// kwhisk hello! world package 'hello!' contains illegal characters
	// michele! hello world namespace 'michele!' contains illegal characters
}

func ExampleFolderManager() {
	SysSh("@rm -Rf /tmp/kwtest")
	fm := NewFolderManager("kwhisk")
	show(fm.ListNamespaces())
	show(fm.ListPackages())
	show(fm.DeletePackage("hello"))
	show(fm.UpdatePackage("hello"))
	show(fm.ListPackages())
	fm.UpdatePackage("world")
	show(fm.ListPackages())
	show(fm.ListActions(nil))
	s := "donotexist"
	show(fm.ListActions(&s))
	gr, err := fm.UpdateAction("hello")
	show(err, gr.dir)
	show(fm.ListActions(nil))
	gr, err = fm.UpdateAction("hello/world")
	show(err, gr.dir)
	show(fm.ListActions(nil))
	fm.UpdateAction("hello/hello")
	fm.UpdatePackage("world")
	fm.UpdateAction("world/world")
	show(fm.ListPackages())
	show(fm.ListActions(nil))
	show(fm.DeletePackage("world"))
	show(fm.DeleteAction("world/hello"))
	show(fm.DeleteAction("world/world"))
	show(fm.DeletePackage("world"))
	show(fm.ListPackages())
	fm.DeleteAction("hello")
	show(fm.ListActions(nil))
	// Output:
	// [kwhisk]
	// []
	// the requested resource does not exists
	// <nil>
	// [hello]
	// [hello world]
	// []
	// []
	// <nil> /tmp/kwtest/kwhisk/default/hello
	// [/kwhisk/default/hello]
	// <nil> /tmp/kwtest/kwhisk/hello/world
	// [/kwhisk/default/hello /kwhisk/hello/world]
	// [hello world]
	// [/kwhisk/default/hello /kwhisk/hello/hello /kwhisk/hello/world /kwhisk/world/world]
	// package not empty
	// the requested resource does not exists
	// <nil>
	// <nil>
	// [hello]
	// [/kwhisk/hello/hello /kwhisk/hello/world]
}
