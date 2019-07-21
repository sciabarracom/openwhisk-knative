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
	return middleware.NotImplemented("operation packages.GetAllPackages has not yet been implemented")
}

// UpdatePackage does it
func UpdatePackage(params packages.UpdatePackageParams, principal *models.Auth) middleware.Responder {
	return middleware.NotImplemented("operation packages.UpdatePackage has not yet been implemented")
}
