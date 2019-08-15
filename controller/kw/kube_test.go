package kw

func ExampleNewKube() {
	k := NewKube()
	grep("tekton", k.ListNamespaces())
	// Output:
	// (string) (len=16) "tekton-pipelines"
}

func ExampleParseK8SYaml() {
	grep(" (Kind|Name):", parseK8sYaml(`---
apiVersion: v1
kind: Namespace
metadata:
  name: knative-whisk
`))
	// Output:
	// Kind: (string) (len=9) "Namespace",
	// Name: (string) (len=13) "knative-whisk",
}
