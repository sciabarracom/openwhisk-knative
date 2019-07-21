package kw

import (
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/models"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/actions"
)

// ConfigureActionsAPI does it
func ConfigureActionsAPI(api *operations.OpenWhiskRESTAPI) {
	// get all
	api.ActionsGetAllActionsHandler = actions.GetAllActionsHandlerFunc(GetAllActions)

	// update
	api.ActionsUpdateActionHandler = actions.UpdateActionHandlerFunc(UpdateAction)
}

// GetAllActions does it
func GetAllActions(params actions.GetAllActionsParams, principal *models.Auth) middleware.Responder {
	return middleware.NotImplemented("operation actions.GetAllActions has not yet been implemented")
}

// UpdateAction does it
func UpdateAction(params actions.UpdateActionParams, principal *models.Auth) middleware.Responder {
	return middleware.NotImplemented("operation actions.UpdateAction has not yet been implemented")
}
