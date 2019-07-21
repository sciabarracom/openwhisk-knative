package kw

func ExampleNewGitRepo() {
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
