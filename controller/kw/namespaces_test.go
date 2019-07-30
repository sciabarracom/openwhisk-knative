package kw

import (
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/namespaces"
)

func ExampleGetAllNamespaces() {
	print(GetAllNamespaces(namespaces.GetAllNamespacesParams{}, nil))
	// Output:
	// &{[kwhisk]}
}
