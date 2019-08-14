package kw

func ExampleNewKube() {
	k := NewKube()
	grep("tekton", k.ListNamespaces())
	// Output:
	// (string) (len=16) "tekton-pipelines"
}
