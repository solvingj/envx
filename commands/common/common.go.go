package common

import (
	"github.com/codegangsta/cli"
	"github.com/sirupsen/logrus"
	"github.com/solvingj/envx/system"
	"os"
)

func PrintHelpAndExitWithError(msg string, context *cli.Context) {
	logrus.Error(msg)
	os.Exit(system.ExitCodeError.Code)
}