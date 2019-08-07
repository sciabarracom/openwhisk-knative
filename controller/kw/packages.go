package kw

import (
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/models"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/restapi/operations/packages"
)

// ConfigurePackagesAPI does it
func ConfigurePackagesAPI(api *operations.OpenWhiskRESTAPI) {
	api.PackagesGetAllPackagesHandler = packages.GetAllPackagesHandlerFunc(GetAllPackages)
	api.PackagesUpdatePackageHandler = packages.UpdatePackageHandlerFunc(UpdatePackage)
	api.PackagesDeletePackageHandler = packages.DeletePackageHandlerFunc(DeletePackage)
}

// GetAllPackages does it
func GetAllPackages(params packages.GetAllPackagesParams, principal *models.Auth) (resp middleware.Responder) {
	defer RecoverRest(&resp)
	payload := []*models.Package{}
	for _, packge := range Manager.ListPackages() {
		//log.Debug(packge)
		s := packge
		payload = append(payload, &models.Package{
			Name:      &s,
			Namespace: &Manager.Namespace,
		})
	}
	return packages.NewGetAllPackagesOK().WithPayload(payload)
}

// UpdatePackage does it
func UpdatePackage(params packages.UpdatePackageParams, principal *models.Auth) (resp middleware.Responder) {
	defer RecoverRest(&resp)
	err := Manager.UpdatePackage(params.PackageName)
	if err != nil {
		return packages.NewUpdatePackageBadRequest().WithPayload(MkErr(err))
	}
	return packages.NewUpdatePackageOK().WithPayload(&models.Package{
		Name:      &params.PackageName,
		Namespace: &Manager.Namespace,
	})
}

// DeletePackage does it
func DeletePackage(params packages.DeletePackageParams, principal *models.Auth) (resp middleware.Responder) {
	defer RecoverRest(&resp)
	err := Manager.DeletePackage(params.PackageName)
	if err != nil {
		return packages.NewDeletePackageConflict().WithPayload(MkErr(err))
	}
	return packages.NewDeletePackageOK()
}
