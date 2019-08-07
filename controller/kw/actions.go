package kw

import (
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/models"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/actions"
)

// ConfigureActionsAPI does it
func ConfigureActionsAPI(api *operations.OpenWhiskRESTAPI) {
	api.ActionsGetAllActionsHandler = actions.GetAllActionsHandlerFunc(GetAllActions)
	api.ActionsUpdateActionHandler = actions.UpdateActionHandlerFunc(UpdateAction)
	api.ActionsUpdateActionInPackageHandler = actions.UpdateActionInPackageHandlerFunc(UpdateActionInPackage)
	api.ActionsDeleteActionHandler = actions.DeleteActionHandlerFunc(DeleteAction)
	api.ActionsDeleteActionInPackageHandler = actions.DeleteActionInPackageHandlerFunc(DeleteActionInPackage)
}

// GetAllActions does it
func GetAllActions(params actions.GetAllActionsParams, principal *models.Auth) (resp middleware.Responder) {
	defer RecoverRest(&resp)
	payload := []*models.Action{}
	for _, action := range Manager.ListActions(nil) {
		ns, pkg, act, _ := Manager.SplitActionName(action)
		if pkg != "default" {
			act = pkg + "/" + act
		}
		payload = append(payload, &models.Action{
			Name:      &act,
			Namespace: &ns,
		})
	}
	return actions.NewGetAllActionsOK().WithPayload(payload)
}

// UpdateAction does it
func UpdateAction(params actions.UpdateActionParams, principal *models.Auth) (resp middleware.Responder) {
	defer RecoverRest(&resp)
	fullname := "/" + params.Namespace + "/" + params.ActionName
	git, err := Manager.UpdateAction(fullname)
	if err != nil {
		return actions.NewUpdateActionBadRequest().WithPayload(MkErr(err))
	}
	err = git.Store("src", []byte(params.Action.Exec.Code))
	if err != nil {
		return actions.NewUpdateActionBadRequest().WithPayload(MkErr(err))
	}
	res := &models.Action{
		Name:      &params.ActionName,
		Namespace: &params.Namespace,
	}
	return actions.NewUpdateActionOK().WithPayload(res)
}

// UpdateActionInPackage does it
func UpdateActionInPackage(params actions.UpdateActionInPackageParams, principal *models.Auth) (resp middleware.Responder) {
	defer RecoverRest(&resp)
	fullname := "/" + params.Namespace +
		"/" + params.PackageName +
		"/" + params.ActionName
	git, err := Manager.UpdateAction(fullname)
	if err != nil {
		return actions.NewUpdateActionInPackageBadRequest().WithPayload(MkErr(err))
	}
	err = git.Store("src", []byte(params.Action.Exec.Code))
	if err != nil {
		return actions.NewUpdateActionInPackageBadRequest().WithPayload(MkErr(err))
	}
	actionName := params.PackageName + "/" + params.ActionName
	res := &models.Action{
		Name:      &actionName,
		Namespace: &params.Namespace,
	}
	return actions.NewUpdateActionInPackageOK().WithPayload(res)
}

// DeleteAction does it
func DeleteAction(params actions.DeleteActionParams, principal *models.Auth) (resp middleware.Responder) {
	defer RecoverRest(&resp)
	err := Manager.DeleteAction(params.ActionName)
	if err != nil {
		return actions.NewDeleteActionConflict().WithPayload(MkErr(err))
	}
	return actions.NewDeleteActionOK()
}

// DeleteActionInPackage does it
func DeleteActionInPackage(params actions.DeleteActionInPackageParams, principal *models.Auth) (resp middleware.Responder) {
	defer RecoverRest(&resp)
	name := params.PackageName + "/" + params.ActionName
	err := Manager.DeleteAction(name)
	if err != nil {
		return actions.NewDeleteActionInPackageConflict().WithPayload(MkErr(err))
	}
	return actions.NewDeleteActionInPackageOK()
}
