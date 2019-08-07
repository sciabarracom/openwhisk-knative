package kw

import (
	"github.com/sciabarracom/openwhisk-knative/controller/gen/models"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/actions"
)

func ExampleGetAllActions() {
	SysSh("@rm -Rf /tmp/kwtest")
	Manager = NewFolderManager("kwhisk")
	params := actions.GetAllActionsParams{}
	show("# None")
	dump(GetAllActions(params, nil))
	Manager.UpdateAction("hello")
	show("# One action")
	grep("Name:", GetAllActions(params, nil))
	show("# Two actions")
	Manager.UpdatePackage("hello")
	Manager.UpdateAction("hello/world")
	grep("Name:", GetAllActions(params, nil))
	show("# Four actions")
	Manager.UpdateAction("world")
	Manager.UpdateAction("hello/hello")
	grep("Name:", GetAllActions(params, nil))
	// Output:
	// # None
	// (*actions.GetAllActionsOK)({
	//  Payload: ([]*models.Action) {
	//  }
	// })
	// # One action
	// Name: (*string)((len=5) "hello"),
	// # Two actions
	// Name: (*string)((len=5) "hello"),
	// Name: (*string)((len=11) "hello/world"),
	// # Four actions
	// Name: (*string)((len=5) "hello"),
	// Name: (*string)((len=5) "world"),
	// Name: (*string)((len=11) "hello/hello"),
	// Name: (*string)((len=11) "hello/world"),
}

func ExampleUpdateAction() {
	SysSh("@rm -Rf /tmp/kwtest")
	Manager = NewFolderManager("kwhisk")
	params := actions.UpdateActionParams{
		ActionName: "bad!",
		Action: &models.ActionPut{
			Exec: &models.ActionExec{
				Code: "action body",
			},
		},
	}
	grep("Error:", UpdateAction(params, nil))
	params.ActionName = "hello"
	params.Namespace = "_"
	grep("Name(space)?:", UpdateAction(params, nil))
	fdump("/tmp/kwtest/kwhisk/default/hello/src")
	// Output:
	// Error: (*string)((len=41) "action 'bad!' contains illegal characters")
	// Name: (*string)((len=5) "hello"),
	// Namespace: (*string)((len=1) "_"),
	// action body
}

func ExampleUpdateActionInPackage() {
	SysSh("@rm -Rf /tmp/kwtest")
	Manager = NewFolderManager("kwhisk")
	params := actions.UpdateActionInPackageParams{
		Namespace:   "_",
		ActionName:  "bad!",
		PackageName: "hello",
		Action: &models.ActionPut{
			Exec: &models.ActionExec{
				Code: "hello world body",
			},
		},
	}
	grep("Error:", UpdateActionInPackage(params, nil))
	params.ActionName = "world"
	grep("Error:", UpdateActionInPackage(params, nil))
	Manager.UpdatePackage("hello")
	grep("Name:", UpdateActionInPackage(params, nil))
	fdump("/tmp/kwtest/kwhisk/hello/world/src")
	// Output:
	// Error: (*string)((len=41) "action 'bad!' contains illegal characters")
	// Error: (*string)((len=43) "package hello in namespace kwhisk not found")
	// Name: (*string)((len=11) "hello/world"),
	// hello world body
}

func ExampleDeleteAction() {
	SysSh("@rm -Rf /tmp/kwtest")
	Manager = NewFolderManager("kwhisk")
	params := actions.DeleteActionParams{
		ActionName: "hello",
	}
	grep("Error:", DeleteAction(params, nil))
	Manager.UpdateAction(("hello"))
	grep("/default/hello", Manager.ListActions(nil))
	dump(DeleteAction(params, nil))
	dump(Manager.ListActions(nil))
	// Output:
	// Error: (*string)((len=38) "the requested resource does not exists")
	// (string) (len=21) "/kwhisk/default/hello"
	// (*actions.DeleteActionOK)({
	// })
	// ([]string) {
	// }
}

func ExampleDeleteActionInPackage() {
	SysSh("@rm -Rf /tmp/kwtest")
	Manager = NewFolderManager("kwhisk")
	params := actions.DeleteActionInPackageParams{
		ActionName:  "world",
		PackageName: "hello",
	}
	grep("Error:", DeleteActionInPackage(params, nil))
	Manager.UpdatePackage("hello")
	Manager.UpdateAction("hello/world")
	grep("hello/world", Manager.ListActions(nil))
	dump(DeleteActionInPackage(params, nil))
	dump(Manager.ListActions(nil))
	// Output:
	// Error: (*string)((len=38) "the requested resource does not exists")
	// (string) (len=19) "/kwhisk/hello/world"
	// (*actions.DeleteActionInPackageOK)({
	// })
	// ([]string) {
	// }
}
