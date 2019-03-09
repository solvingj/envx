package main

import (
	"github.com/sirupsen/logrus"
	"github.com/solvingj/envx/commands"
	"github.com/solvingj/envx/system"
	"os"
)

func main() {
	var log = logrus.New()
	log.SetLevel(GetCliLogLevel())
	err := execMain()
	system.ExitOnErr(err)
}

func execMain() error {
	app := commands.GetApp()
	args := os.Args
	err := app.Run(args)
	return err
}

func GetCliLogLevel() logrus.Level {
	switch os.Getenv("ENVX_CLI_LOG_LEVEL") {
	case "ERROR":
		return logrus.ErrorLevel
	case "WARN":
		return logrus.WarnLevel
	case "DEBUG":
		return logrus.DebugLevel
	default:
		return logrus.InfoLevel
	}
}