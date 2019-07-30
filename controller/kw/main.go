package kw

import (
	"os"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"

	//runtime "github.com/go-openapi/runtime"

	"github.com/go-openapi/loads"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/models"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/actions"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/activations"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/packages"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/rules"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/triggers"
)

// Manager manages entities (as folders)
var Manager *FolderManager

func init() {
	// managers
	defaultNamespace := os.Getenv("KW_NAMESPACE")
	if defaultNamespace == "" {
		defaultNamespace = "kwhisk"
	}
	Manager = NewFolderManager(defaultNamespace)
}

// Main starts the server
func Main() {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	PanicIf(err)

	// configure APIs
	api := operations.NewOpenWhiskRESTAPI(swaggerSpec)
	api.BasicAuthAuth = BasicAuth
	ConfigureNamespacesAPI(api)
	ConfigurePackagesAPI(api)
	ConfigureActionsAPI(api)
	ConfigureUnimplementedAPI(api)

	// server
	server := restapi.NewServer(api)
	defer server.Shutdown()
	server.Port = 8080
	PanicIf(server.Serve())
}

// BasicAuth performs basic authentication
func BasicAuth(user string, pass string) (*models.Auth, error) {
	if user == "123" && pass == "456" {
		return &models.Auth{Username: user, Password: pass}, nil
	}
	return nil, errors.Unauthenticated("bad username or password")
}

// ConfigureUnimplementedAPI does it
func ConfigureUnimplementedAPI(api *operations.OpenWhiskRESTAPI) {

	// unimplemented methods
	api.ActionsDeleteWebNamespacePackageNameActionNameExtensionHandler = actions.DeleteWebNamespacePackageNameActionNameExtensionHandlerFunc(func(params actions.DeleteWebNamespacePackageNameActionNameExtensionParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation actions.DeleteWebNamespacePackageNameActionNameExtension has not yet been implemented")
	})
	api.ActionsGetWebNamespacePackageNameActionNameExtensionHandler = actions.GetWebNamespacePackageNameActionNameExtensionHandlerFunc(func(params actions.GetWebNamespacePackageNameActionNameExtensionParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation actions.GetWebNamespacePackageNameActionNameExtension has not yet been implemented")
	})
	api.ActionsPostWebNamespacePackageNameActionNameExtensionHandler = actions.PostWebNamespacePackageNameActionNameExtensionHandlerFunc(func(params actions.PostWebNamespacePackageNameActionNameExtensionParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation actions.PostWebNamespacePackageNameActionNameExtension has not yet been implemented")
	})
	api.ActionsPutWebNamespacePackageNameActionNameExtensionHandler = actions.PutWebNamespacePackageNameActionNameExtensionHandlerFunc(func(params actions.PutWebNamespacePackageNameActionNameExtensionParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation actions.PutWebNamespacePackageNameActionNameExtension has not yet been implemented")
	})
	api.ActionsDeleteActionHandler = actions.DeleteActionHandlerFunc(func(params actions.DeleteActionParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation actions.DeleteAction has not yet been implemented")
	})
	api.ActionsDeleteActionInPackageHandler = actions.DeleteActionInPackageHandlerFunc(func(params actions.DeleteActionInPackageParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation actions.DeleteActionInPackage has not yet been implemented")
	})
	api.PackagesDeletePackageHandler = packages.DeletePackageHandlerFunc(func(params packages.DeletePackageParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation packages.DeletePackage has not yet been implemented")
	})

	api.RulesDeleteRuleHandler = rules.DeleteRuleHandlerFunc(func(params rules.DeleteRuleParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation rules.DeleteRule has not yet been implemented")
	})

	api.TriggersDeleteTriggerHandler = triggers.DeleteTriggerHandlerFunc(func(params triggers.DeleteTriggerParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation triggers.DeleteTrigger has not yet been implemented")
	})
	api.TriggersFireTriggerHandler = triggers.FireTriggerHandlerFunc(func(params triggers.FireTriggerParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation triggers.FireTrigger has not yet been implemented")
	})
	api.ActionsGetActionByNameHandler = actions.GetActionByNameHandlerFunc(func(params actions.GetActionByNameParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation actions.GetActionByName has not yet been implemented")
	})
	api.ActionsGetActionInPackageByNameHandler = actions.GetActionInPackageByNameHandlerFunc(func(params actions.GetActionInPackageByNameParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation actions.GetActionInPackageByName has not yet been implemented")
	})
	api.ActivationsGetActivationByIDHandler = activations.GetActivationByIDHandlerFunc(func(params activations.GetActivationByIDParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation activations.GetActivationByID has not yet been implemented")
	})
	api.ActivationsGetActivationLogsHandler = activations.GetActivationLogsHandlerFunc(func(params activations.GetActivationLogsParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation activations.GetActivationLogs has not yet been implemented")
	})
	api.ActivationsGetActivationResultHandler = activations.GetActivationResultHandlerFunc(func(params activations.GetActivationResultParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation activations.GetActivationResult has not yet been implemented")
	})
	api.ActivationsGetActivationsHandler = activations.GetActivationsHandlerFunc(func(params activations.GetActivationsParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation activations.GetActivations has not yet been implemented")
	})
	api.RulesGetAllRulesHandler = rules.GetAllRulesHandlerFunc(func(params rules.GetAllRulesParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation rules.GetAllRules has not yet been implemented")
	})
	api.TriggersGetAllTriggersHandler = triggers.GetAllTriggersHandlerFunc(func(params triggers.GetAllTriggersParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation triggers.GetAllTriggers has not yet been implemented")
	})
	api.PackagesGetPackageByNameHandler = packages.GetPackageByNameHandlerFunc(func(params packages.GetPackageByNameParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation packages.GetPackageByName has not yet been implemented")
	})
	api.RulesGetRuleByNameHandler = rules.GetRuleByNameHandlerFunc(func(params rules.GetRuleByNameParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation rules.GetRuleByName has not yet been implemented")
	})
	api.TriggersGetTriggerByNameHandler = triggers.GetTriggerByNameHandlerFunc(func(params triggers.GetTriggerByNameParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation triggers.GetTriggerByName has not yet been implemented")
	})
	api.ActionsInvokeActionHandler = actions.InvokeActionHandlerFunc(func(params actions.InvokeActionParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation actions.InvokeAction has not yet been implemented")
	})
	api.ActionsInvokeActionInPackageHandler = actions.InvokeActionInPackageHandlerFunc(func(params actions.InvokeActionInPackageParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation actions.InvokeActionInPackage has not yet been implemented")
	})
	api.RulesSetStateHandler = rules.SetStateHandlerFunc(func(params rules.SetStateParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation rules.SetState has not yet been implemented")
	})
	api.ActionsUpdateActionInPackageHandler = actions.UpdateActionInPackageHandlerFunc(func(params actions.UpdateActionInPackageParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation actions.UpdateActionInPackage has not yet been implemented")
	})
	api.RulesUpdateRuleHandler = rules.UpdateRuleHandlerFunc(func(params rules.UpdateRuleParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation rules.UpdateRule has not yet been implemented")
	})
	api.TriggersUpdateTriggerHandler = triggers.UpdateTriggerHandlerFunc(func(params triggers.UpdateTriggerParams, principal *models.Auth) middleware.Responder {
		return middleware.NotImplemented("operation triggers.UpdateTrigger has not yet been implemented")
	})
}
