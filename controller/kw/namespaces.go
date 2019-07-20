package kw

import (
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/models"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/namespaces"
)

// GetAllNamespaces does it
func GetAllNamespaces(params namespaces.GetAllNamespacesParams, principal *models.Auth) middleware.Responder {
	return namespaces.NewGetAllNamespacesOK().WithPayload([]string{"knative"})
}
