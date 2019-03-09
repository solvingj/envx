package commands

import (
	"github.com/codegangsta/cli"
	"github.com/solvingj/envx/commands/config"
	"github.com/solvingj/envx/commands/env"
	"github.com/solvingj/envx/commands/run"
	docs "github.com/solvingj/envx/docs/envx"
	configdocs "github.com/solvingj/envx/docs/envx/config"
	envdocs "github.com/solvingj/envx/docs/envx/env"
	rundocs "github.com/solvingj/envx/docs/envx/run"
)

const version = "0.1.0"

func GetApp() *cli.App {
	app := cli.NewApp()
	app.Name = "envx"
	app.Version = version
	app.Usage = docs.AppDescription
	app.Commands = GetCommands()
	app.HideHelp = true
	app.Flags = GlobalOptions()
	return app
}

func GlobalOptions() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "log-level",
			Usage: docs.LogLevelUsage,
			EnvVar: "ENVX_LOG_LEVEL",

		},
		cli.StringFlag{
			Name:  "envx-home",
			Usage: docs.HomeDirectoryUsage,
			EnvVar: "ENVX_HOME_DIR",
		},
	}
}


func GetCommands() []cli.Command {
	return []cli.Command{
		ConfigCommand(),
		EnvCommand(),
		RunCommand(),
	}
}

func EnvCommand() cli.Command {
	return cli.Command{
		Name:        "env",
		Usage:       envdocs.Description,
		Subcommands: env.GetCommands(),
		HideHelp:    true,
	}
}

func ConfigCommand() cli.Command {
	return cli.Command{
		Name:        "config",
		Usage:       configdocs.Description,
		Subcommands: config.GetCommands(),
		HideHelp:    true,
	}
}

func RunCommand() cli.Command {
	return cli.Command{
		Name:      "run",
		Flags:     run.Flags(),
		Usage:     rundocs.Description,
		ArgsUsage: "<command_to_run>",
		HideHelp:  true,
		Action: func(c *cli.Context) {
			run.RunCmd(c)
		},
	}
}


