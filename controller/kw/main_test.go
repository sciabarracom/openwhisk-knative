package kw

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

var Debug = true

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags

	log.SetOutput(os.Stderr)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: true,
	})
	log.SetLevel(log.DebugLevel)
	//log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func trace() {
	log.SetLevel(log.TraceLevel)
}
