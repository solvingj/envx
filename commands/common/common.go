package common

import (
	"github.com/codegangsta/cli"
	"github.com/jfrog/jfrog-client-go/utils/log"
	"os"
	"strings"
)

func CreateArgsUsage(args map[string]string) string {
	var sb strings.Builder
	for k, v := range args{
		sb.WriteString("    " + k + "\n")
		sb.WriteString("    " + k + "\n")
		sb.WriteString("        " + v + "\n")
	}
	return sb.String()
}

func PrintHelpAndExitWithError(msg string, context *cli.Context) {
	log.Error(msg)
	cli.ShowCommandHelp(context, context.Command.Name)
	os.Exit(ExitCodeError.Code)
}


type ExitCode struct {
	Code int
}

var ExitCodeNoError = ExitCode{0}
var ExitCodeError = ExitCode{1}
var ExitCodeFailNoOp = ExitCode{2}
var ExitCodeBuildScan = ExitCode{3}
