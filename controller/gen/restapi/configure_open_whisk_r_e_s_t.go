// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/actions"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/activations"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/namespaces"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/packages"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/rules"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/triggers"

	models "github.com/sciabarracom/openwhisk-knative/controller/gen/models"
)

//go:generate swagger generate server --target ../../gen --name OpenWhiskREST --spec ../../apiv1swagger.json --principal models.Auth --exclude-main

func configureFlags(api *operations.OpenWhiskRESTAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.OpenWhiskRESTAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.TxtProducer = runtime.TextProducer()

	// Applies when the Authorization header is set with the Basic scheme
	api.BasicAuthAuth = func(user string, pass string) (*models.Auth, error) {
		return nil, errors.NotImplemented("basic auth  (basicAuth) has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	if api.ActionsDeleteWebNamespacePackageNameActionNameExtensionHandler == nil {
		api.ActionsDeleteWebNamespacePackageNameActionNameExtensionHandler = actions.DeleteWebNamespacePackageNameActionNameExtensionHandlerFunc(func(params actions.DeleteWebNamespacePackageNameActionNameExtensionParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation actions.DeleteWebNamespacePackageNameActionNameExtension has not yet been implemented")
		})
	}
	if api.ActionsGetWebNamespacePackageNameActionNameExtensionHandler == nil {
		api.ActionsGetWebNamespacePackageNameActionNameExtensionHandler = actions.GetWebNamespacePackageNameActionNameExtensionHandlerFunc(func(params actions.GetWebNamespacePackageNameActionNameExtensionParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation actions.GetWebNamespacePackageNameActionNameExtension has not yet been implemented")
		})
	}
	if api.ActionsPostWebNamespacePackageNameActionNameExtensionHandler == nil {
		api.ActionsPostWebNamespacePackageNameActionNameExtensionHandler = actions.PostWebNamespacePackageNameActionNameExtensionHandlerFunc(func(params actions.PostWebNamespacePackageNameActionNameExtensionParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation actions.PostWebNamespacePackageNameActionNameExtension has not yet been implemented")
		})
	}
	if api.ActionsPutWebNamespacePackageNameActionNameExtensionHandler == nil {
		api.ActionsPutWebNamespacePackageNameActionNameExtensionHandler = actions.PutWebNamespacePackageNameActionNameExtensionHandlerFunc(func(params actions.PutWebNamespacePackageNameActionNameExtensionParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation actions.PutWebNamespacePackageNameActionNameExtension has not yet been implemented")
		})
	}
	if api.ActionsDeleteActionHandler == nil {
		api.ActionsDeleteActionHandler = actions.DeleteActionHandlerFunc(func(params actions.DeleteActionParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation actions.DeleteAction has not yet been implemented")
		})
	}
	if api.ActionsDeleteActionInPackageHandler == nil {
		api.ActionsDeleteActionInPackageHandler = actions.DeleteActionInPackageHandlerFunc(func(params actions.DeleteActionInPackageParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation actions.DeleteActionInPackage has not yet been implemented")
		})
	}
	if api.PackagesDeletePackageHandler == nil {
		api.PackagesDeletePackageHandler = packages.DeletePackageHandlerFunc(func(params packages.DeletePackageParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation packages.DeletePackage has not yet been implemented")
		})
	}
	if api.RulesDeleteRuleHandler == nil {
		api.RulesDeleteRuleHandler = rules.DeleteRuleHandlerFunc(func(params rules.DeleteRuleParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation rules.DeleteRule has not yet been implemented")
		})
	}
	if api.TriggersDeleteTriggerHandler == nil {
		api.TriggersDeleteTriggerHandler = triggers.DeleteTriggerHandlerFunc(func(params triggers.DeleteTriggerParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation triggers.DeleteTrigger has not yet been implemented")
		})
	}
	if api.TriggersFireTriggerHandler == nil {
		api.TriggersFireTriggerHandler = triggers.FireTriggerHandlerFunc(func(params triggers.FireTriggerParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation triggers.FireTrigger has not yet been implemented")
		})
	}
	if api.ActionsGetActionByNameHandler == nil {
		api.ActionsGetActionByNameHandler = actions.GetActionByNameHandlerFunc(func(params actions.GetActionByNameParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation actions.GetActionByName has not yet been implemented")
		})
	}
	if api.ActionsGetActionInPackageByNameHandler == nil {
		api.ActionsGetActionInPackageByNameHandler = actions.GetActionInPackageByNameHandlerFunc(func(params actions.GetActionInPackageByNameParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation actions.GetActionInPackageByName has not yet been implemented")
		})
	}
	if api.ActivationsGetActivationByIDHandler == nil {
		api.ActivationsGetActivationByIDHandler = activations.GetActivationByIDHandlerFunc(func(params activations.GetActivationByIDParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation activations.GetActivationByID has not yet been implemented")
		})
	}
	if api.ActivationsGetActivationLogsHandler == nil {
		api.ActivationsGetActivationLogsHandler = activations.GetActivationLogsHandlerFunc(func(params activations.GetActivationLogsParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation activations.GetActivationLogs has not yet been implemented")
		})
	}
	if api.ActivationsGetActivationResultHandler == nil {
		api.ActivationsGetActivationResultHandler = activations.GetActivationResultHandlerFunc(func(params activations.GetActivationResultParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation activations.GetActivationResult has not yet been implemented")
		})
	}
	if api.ActivationsGetActivationsHandler == nil {
		api.ActivationsGetActivationsHandler = activations.GetActivationsHandlerFunc(func(params activations.GetActivationsParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation activations.GetActivations has not yet been implemented")
		})
	}
	if api.ActionsGetAllActionsHandler == nil {
		api.ActionsGetAllActionsHandler = actions.GetAllActionsHandlerFunc(func(params actions.GetAllActionsParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation actions.GetAllActions has not yet been implemented")
		})
	}
	if api.NamespacesGetAllNamespacesHandler == nil {
		api.NamespacesGetAllNamespacesHandler = namespaces.GetAllNamespacesHandlerFunc(func(params namespaces.GetAllNamespacesParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation namespaces.GetAllNamespaces has not yet been implemented")
		})
	}
	if api.PackagesGetAllPackagesHandler == nil {
		api.PackagesGetAllPackagesHandler = packages.GetAllPackagesHandlerFunc(func(params packages.GetAllPackagesParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation packages.GetAllPackages has not yet been implemented")
		})
	}
	if api.RulesGetAllRulesHandler == nil {
		api.RulesGetAllRulesHandler = rules.GetAllRulesHandlerFunc(func(params rules.GetAllRulesParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation rules.GetAllRules has not yet been implemented")
		})
	}
	if api.TriggersGetAllTriggersHandler == nil {
		api.TriggersGetAllTriggersHandler = triggers.GetAllTriggersHandlerFunc(func(params triggers.GetAllTriggersParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation triggers.GetAllTriggers has not yet been implemented")
		})
	}
	if api.PackagesGetPackageByNameHandler == nil {
		api.PackagesGetPackageByNameHandler = packages.GetPackageByNameHandlerFunc(func(params packages.GetPackageByNameParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation packages.GetPackageByName has not yet been implemented")
		})
	}
	if api.RulesGetRuleByNameHandler == nil {
		api.RulesGetRuleByNameHandler = rules.GetRuleByNameHandlerFunc(func(params rules.GetRuleByNameParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation rules.GetRuleByName has not yet been implemented")
		})
	}
	if api.TriggersGetTriggerByNameHandler == nil {
		api.TriggersGetTriggerByNameHandler = triggers.GetTriggerByNameHandlerFunc(func(params triggers.GetTriggerByNameParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation triggers.GetTriggerByName has not yet been implemented")
		})
	}
	if api.ActionsInvokeActionHandler == nil {
		api.ActionsInvokeActionHandler = actions.InvokeActionHandlerFunc(func(params actions.InvokeActionParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation actions.InvokeAction has not yet been implemented")
		})
	}
	if api.ActionsInvokeActionInPackageHandler == nil {
		api.ActionsInvokeActionInPackageHandler = actions.InvokeActionInPackageHandlerFunc(func(params actions.InvokeActionInPackageParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation actions.InvokeActionInPackage has not yet been implemented")
		})
	}
	if api.RulesSetStateHandler == nil {
		api.RulesSetStateHandler = rules.SetStateHandlerFunc(func(params rules.SetStateParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation rules.SetState has not yet been implemented")
		})
	}
	if api.ActionsUpdateActionHandler == nil {
		api.ActionsUpdateActionHandler = actions.UpdateActionHandlerFunc(func(params actions.UpdateActionParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation actions.UpdateAction has not yet been implemented")
		})
	}
	if api.ActionsUpdateActionInPackageHandler == nil {
		api.ActionsUpdateActionInPackageHandler = actions.UpdateActionInPackageHandlerFunc(func(params actions.UpdateActionInPackageParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation actions.UpdateActionInPackage has not yet been implemented")
		})
	}
	if api.PackagesUpdatePackageHandler == nil {
		api.PackagesUpdatePackageHandler = packages.UpdatePackageHandlerFunc(func(params packages.UpdatePackageParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation packages.UpdatePackage has not yet been implemented")
		})
	}
	if api.RulesUpdateRuleHandler == nil {
		api.RulesUpdateRuleHandler = rules.UpdateRuleHandlerFunc(func(params rules.UpdateRuleParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation rules.UpdateRule has not yet been implemented")
		})
	}
	if api.TriggersUpdateTriggerHandler == nil {
		api.TriggersUpdateTriggerHandler = triggers.UpdateTriggerHandlerFunc(func(params triggers.UpdateTriggerParams, principal *models.Auth) middleware.Responder {
			return middleware.NotImplemented("operation triggers.UpdateTrigger has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
