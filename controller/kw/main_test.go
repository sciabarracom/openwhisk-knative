package kw

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	Testing = true
	Debugging = false
	os.Exit(m.Run())
}
