package kw

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/prometheus/common/log"
	"github.com/sirupsen/logrus"
)

// Debugging ?
var Debugging = false

// Testing ?
var Testing = false

// Debug messages if debugging
func Debug(args ...interface{}) {
	if Debugging {
		log.Debug(args...)
	}
}

// PanicIf panics if error is not nil
func PanicIf(err error) {
	if err != nil {
		logrus.Error(err)
		if Testing {
			fmt.Printf("ERR: %s\n", err.Error())
		}
		panic(err)
	}
}

// LogIf logs a warning if the error is not nil
func LogIf(err error) {
	if err != nil {
		logrus.Warn(err)
		if Testing {
			fmt.Printf("WARN: %s\n", err.Error())
		}
	}
}

// Recover recovers from a panic returning an error
func Recover(perr *error) {
	if r := recover(); r != nil {
		switch v := r.(type) {
		case error:
			*perr = v
		case string:
			*perr = fmt.Errorf("%s", v)
		default:
			*perr = fmt.Errorf("%v", v)
		}
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
