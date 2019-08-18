package kw

func ExampleNewKube() {
	k := NewKube()
	grep("tekton", k.ListNamespaces())
	// Output:
	// (string) (len=16) "tekton-pipelines"
}

func ExampleParseK8SYaml() {
	//
	o := parseK8sYaml(`---
apiVersion: v1
kind: Namespace
metadata:
  name: knative-whisk
`)
	//show(o[0].(*v1.Namespace))
	grep(" (Kind|Name):", o)
	// Output:
	// Kind: (string) (len=9) "Namespace",
	// Name: (string) (len=13) "knative-whisk",
}

func ExampleBuild() {
	s := build("alpha", "beta", "gamma")
	show(s)
	o := parseK8sYaml(s)
	dump(o)
	// Output:
	// -
}

func ExampleApply() {
	k := NewKube()
	//grep(" (Kind|Name):",
	k.Apply(`---
apiVersion: v1
kind: Namespace
metadata:
  name: knative-whisk
`)

	// Output:
	// -
}
