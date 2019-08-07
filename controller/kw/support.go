package kw

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/sciabarracom/openwhisk-knative/controller/gen/models"
)

// ValidateURLPathComponent checks if a string is a valid path component
func ValidateURLPathComponent(urlComponent string) bool {
	var check = regexp.MustCompile(`^[[:word:]-.~]+$`)
	return check.MatchString(urlComponent)
}

// MkErr builds an error
func MkErr(err error) *models.ErrorMessage {
	msg := err.Error()
	hash := fmt.Sprintf("%x", md5.Sum([]byte(msg)))
	return &models.ErrorMessage{
		Code:  hash,
		Error: &msg,
	}
}

// RecoverRest recovers from a panic returning a rest error response
func RecoverRest(resp *middleware.Responder) {
	if r := recover(); r != nil {
		switch v := r.(type) {
		case error:
			*resp = &panicResp{v.Error()}
		default:
			*resp = &panicResp{r}
		}
	}
}

/// support code for RecoverRest
type panicResp struct {
	response interface{}
}

func (p *panicResp) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(http.StatusInternalServerError)
	LogIf(producer.Produce(rw, p.response))
}

// IsDirEmpty checks if a directory is empty
func IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// read in ONLY one file
	_, err = f.Readdir(1)

	// and if the file is EOF... well, the dir is empty.
	if err == io.EOF {
		return true, nil
	}
	return false, err
}
