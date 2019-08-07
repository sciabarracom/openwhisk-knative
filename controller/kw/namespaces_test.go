package kw

import (
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/namespaces"
)

func ExampleGetAllNamespaces() {
	show(GetAllNamespaces(namespaces.GetAllNamespacesParams{}, nil))
	// Output:
	// &{[kwhisk]}
}
