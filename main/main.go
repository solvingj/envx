package main

import (
	"github.com/jfrog/jfrog-cli-go/jfrog-cli/utils/cliutils"
	"github.com/jfrog/jfrog-client-go/utils/log"
	"github.com/solvingj/envx/commands"
	"os"
)

func main() {
	log.SetLogger(log.NewLogger(GetCliLogLevel()))
	err := execMain()
	cliutils.ExitOnErr(err)
}

func execMain() error {
	app := commands.GetApp()
	args := os.Args
	err := app.Run(args)
	return err
}

func GetCliLogLevel() log.LevelType {
	switch os.Getenv("ENVX_CLI_LOG_LEVEL") {
	case "ERROR":
		return log.ERROR
	case "WARN":
		return log.WARN
	case "DEBUG":
		return log.DEBUG
	default:
		return log.INFO
	}
}