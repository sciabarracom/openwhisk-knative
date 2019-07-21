package kw

import (
	"fmt"

	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/namespaces"
)

func ExampleGetAllNamespaces() {
	fmt.Print(GetAllNamespaces(namespaces.GetAllNamespacesParams{}, nil))
	// Output:
	// &{[knative]}
}
