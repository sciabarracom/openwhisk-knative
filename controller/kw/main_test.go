package kw

import (
	"io/ioutil"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

var Debug = true

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	if Debug {
		log.SetOutput(os.Stderr)
		log.SetFormatter(&log.TextFormatter{
			DisableColors:    true,
			DisableTimestamp: true,
		})
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetOutput(ioutil.Discard)
	}
	os.Exit(m.Run())
}
