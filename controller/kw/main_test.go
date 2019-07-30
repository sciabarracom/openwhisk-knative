package kw

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

// ConfigState for spew
var sp spew.ConfigState = spew.ConfigState{
	Indent:                  " ",
	DisablePointerAddresses: true,
	DisableCapacities:       true,
	SortKeys:                true,
	SpewKeys:                true,
	DisableMethods:          true,
}

func print(args ...interface{}) {
	fmt.Println(args...)
}

func printf(format string, args ...interface{}) {
	fmt.Printf(format+"\\n", args...)
}

func dump(args ...interface{}) {
	sp.Dump(args...)
}

func grep(search string, data ...interface{}) {
	re := regexp.MustCompile(search)
	lines := strings.Split(sp.Sdump(data...), "\n")
	for _, line := range lines {
		//print(line)
		if re.Match([]byte(line)) {
			print(strings.TrimSpace(line))
		}
	}
}

type recovering func()

func capture(fn recovering) {
	defer func() {
		if r := recover(); r != nil {
			print("capture:", r)
		}
	}()
	fn()
}

func debug(arg ...interface{}) {
	log.Debug(arg...)
}

func trace(arg ...interface{}) {
	log.Trace(arg...)
}

func traceOn() {
	log.SetLevel(log.TraceLevel)
}

func traceOff() {
	log.SetLevel(log.DebugLevel)
}

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
