package system

import (
	log "github.com/sirupsen/logrus"
	"os"
)

// Exit codes:
type ExitCode struct {
	Code int
}

var ExitCodeNoError = ExitCode{0}
var ExitCodeError = ExitCode{1}
var ExitCodeFailNoOp = ExitCode{2}
var ExitCodeBuildScan = ExitCode{3}

func PanicOnError(err error) error {
	if err != nil {
		panic(err)
	}
	return err
}

func ExitOnErr(err error) {
	if exitCode := GetExitCode(err, 0, 0, false); exitCode != ExitCodeNoError {
		traceExit(exitCode, err)
	}
}

func FailNoOp(err error, success, failed int, failNoOp bool) {
	if exitCode := GetExitCode(err, success, failed, failNoOp); exitCode != ExitCodeNoError {
		traceExit(exitCode, err)
	}
}

func GetExitCode(err error, success, failed int, failNoOp bool) ExitCode {
	// Error occurred - Return 1
	if err != nil || failed > 0 {
		return ExitCodeError
	}
	// No errors, but also no files affected - Return 2 if failNoOp
	if success == 0 && failNoOp {
		return ExitCodeFailNoOp
	}
	// Otherwise - Return 0
	return ExitCodeNoError
}

func traceExit(exitCode ExitCode, err error) {
	if err != nil && len(err.Error()) > 0 {
		log.Error(err)
	}
	os.Exit(exitCode.Code)
}
