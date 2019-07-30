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
}

// GetAllPackages does it
func GetAllPackages(params packages.GetAllPackagesParams, principal *models.Auth) middleware.Responder {
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
func UpdatePackage(params packages.UpdatePackageParams, principal *models.Auth) middleware.Responder {
	err := Manager.UpdatePackage(params.PackageName)
	if err != nil {
		return packages.NewUpdatePackageBadRequest().WithPayload(MkErr(err))
	}
	return packages.NewUpdatePackageOK().WithPayload(&models.Package{
		Name:      &params.PackageName,
		Namespace: &Manager.Namespace,
	})
}
