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
	/*payload := []*models.Package{}
	for _, packge := range Manager.ListPackages() {
		log.Debug(packge)
		s := packge
		payload = append(payload, &models.Package{
			Name:      &s,
			Namespace: &Manager.Namespace,
		})
	}
	return packages.NewGetAllPackagesOK().WithPayload(payload)*/
}

// UpdateAction does it
func UpdateAction(params actions.UpdateActionParams, principal *models.Auth) middleware.Responder {
	/*err := Manager.UpdateAction(params.ActionName)
	if err != nil {
		return packages.NewUpdatePackageBadRequest().WithPayload(MkErr(err))
	}
	return packages.NewUpdatePackageOK().WithPayload(&models.Package{
		Name:      &params.PackageName,
		Namespace: &Manager.Namespace,
	})*/
	return middleware.NotImplemented("operation actions.UpdateAction has not yet been implemented")
}
