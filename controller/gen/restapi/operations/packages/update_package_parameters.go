// Code generated by go-swagger; DO NOT EDIT.

package packages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sciabarracom/openwhisk-knative/controller/gen/models"
)

// NewUpdatePackageParams creates a new UpdatePackageParams object
// no default values defined in spec.
func NewUpdatePackageParams() UpdatePackageParams {

	return UpdatePackageParams{}
}

// UpdatePackageParams contains all the bound params for the update package operation
// typically these are obtained from a http.Request
//
// swagger:parameters updatePackage
type UpdatePackageParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The entity namespace
	  Required: true
	  In: path
	*/
	Namespace string
	/*Overwrite item if it exists. Default is false.
	  In: query
	*/
	Overwrite *string
	/*The package being updated
	  Required: true
	  In: body
	*/
	Package *models.PackagePut
	/*Name of package
	  Required: true
	  In: path
	*/
	PackageName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdatePackageParams() beforehand.
func (o *UpdatePackageParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rNamespace, rhkNamespace, _ := route.Params.GetOK("namespace")
	if err := o.bindNamespace(rNamespace, rhkNamespace, route.Formats); err != nil {
		res = append(res, err)
	}

	qOverwrite, qhkOverwrite, _ := qs.GetOK("overwrite")
	if err := o.bindOverwrite(qOverwrite, qhkOverwrite, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PackagePut
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("package", "body"))
			} else {
				res = append(res, errors.NewParseError("package", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Package = &body
			}
		}
	} else {
		res = append(res, errors.Required("package", "body"))
	}
	rPackageName, rhkPackageName, _ := route.Params.GetOK("packageName")
	if err := o.bindPackageName(rPackageName, rhkPackageName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindNamespace binds and validates parameter Namespace from path.
func (o *UpdatePackageParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Namespace = raw

	return nil
}

// bindOverwrite binds and validates parameter Overwrite from query.
func (o *UpdatePackageParams) bindOverwrite(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Overwrite = &raw

	if err := o.validateOverwrite(formats); err != nil {
		return err
	}

	return nil
}

// validateOverwrite carries on validations for parameter Overwrite
func (o *UpdatePackageParams) validateOverwrite(formats strfmt.Registry) error {

	if err := validate.Enum("overwrite", "query", *o.Overwrite, []interface{}{"true", "false"}); err != nil {
		return err
	}

	return nil
}

// bindPackageName binds and validates parameter PackageName from path.
func (o *UpdatePackageParams) bindPackageName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.PackageName = raw

	return nil
}