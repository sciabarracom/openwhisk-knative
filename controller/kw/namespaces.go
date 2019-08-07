package kw

import (
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/models"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/namespaces"
)

// ConfigureNamespacesAPI does it
func ConfigureNamespacesAPI(api *operations.OpenWhiskRESTAPI) {
	api.NamespacesGetAllNamespacesHandler = namespaces.GetAllNamespacesHandlerFunc(GetAllNamespaces)
}

// GetAllNamespaces does it
func GetAllNamespaces(params namespaces.GetAllNamespacesParams, principal *models.Auth) (resp middleware.Responder) {
	defer RecoverRest(&resp)
	return namespaces.NewGetAllNamespacesOK().WithPayload(Manager.ListNamespaces())
}
