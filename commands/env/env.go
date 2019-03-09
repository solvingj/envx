package env

import (
	"errors"
	"github.com/codegangsta/cli"
	"github.com/solvingj/envx/commands/common"
	envdocs "github.com/solvingj/envx/docs/envx/env"
	"github.com/solvingj/envx/env"
	"github.com/solvingj/envx/system"
)

func GetCommands() []cli.Command {
	return []cli.Command{
		{
			Name:      "list",
			Usage:     envdocs.DescriptionList,
			HideHelp:  true,
			Action: func(c *cli.Context) {
				ListEnvCmd(c)
			},
		},
		{
			Name:      "show",
			Usage:     envdocs.DescriptionShow,
			ArgsUsage: "<env_name>",
			HideHelp:  true,
			Action: func(c *cli.Context) {
				ShowEnvCmd(c)
			},
		},
		{
			Name:      "new",
			Usage:     envdocs.DescriptionNew,
			ArgsUsage: "<env_name>",
			HideHelp:  true,
			Action: func(c *cli.Context) {
				NewEnvCmd(c)
			},
		},
		{
			Name:      "delete",
			Usage:     envdocs.DescriptionDelete,
			ArgsUsage: "<env_name>",
			HideHelp:  true,
			Action: func(c *cli.Context) {
				DeleteEnvCmd(c)
			},
		},
		{
			Name:      "update",
			Usage:     envdocs.DescriptionUpdate,
			ArgsUsage: "<env_name> <variable=value>",
			HideHelp:  true,
			Action: func(c *cli.Context) {
				UpdateEnvCmd(c)
			},
		},
	}
}

func ListEnvCmd(c *cli.Context) {
	err := env.ListEnvs()
	if err != nil {
		system.ExitOnErr(err)
	}
}

func ShowEnvCmd(c *cli.Context) {
	if len(c.Args()) < 1 {
		common.PrintHelpAndExitWithError("Wrong number of arguments.", c)
	}
	var envId = c.Args().Get(0)
	err := env.ShowEnv(envId)
	if err != nil {
		system.ExitOnErr(err)
	}
}

//TODO: Implement NewEnvCmd function
func NewEnvCmd(c *cli.Context) {
	err := errors.New("command not implemented")
	if err != nil {
		system.ExitOnErr(err)
	}
	//if len(c.Args()) < 4 {
	//	cliutils.PrintHelpAndExitWithError("Wrong number of arguments.", c)
	//} else {
	//	var envId = c.Args()[3]
	//	err := env.New(envId)
	//	if err != nil {
	//		system.ExitOnErr(err)
	//	}
	//}
}

//TODO: Implement DeleteEnvCmd function
func DeleteEnvCmd(c *cli.Context) {
	err := errors.New("command not implemented")
	if err != nil {
		system.ExitOnErr(err)
	}
	//if len(c.Args()) < 4 {
	//	cliutils.PrintHelpAndExitWithError("Wrong number of arguments.", c)
	//} else {
	//	var envId = c.Args()[3]
	//	err := env.New(envId)
	//	if err != nil {
	//		system.ExitOnErr(err)
	//	}
	//}
}

//TODO: Implement DeleteUpdateCmd function
func UpdateEnvCmd(c *cli.Context) {
	err := errors.New("command not implemented")
	if err != nil {
		system.ExitOnErr(err)
	}
	//if len(c.Args()) < 6 {
	//	cliutils.PrintHelpAndExitWithError("Wrong number of arguments.", c)
	//} else {
	//	var envId = c.Args()[3]
	//	err := env.New(envId)
	//	if err != nil {
	//		system.ExitOnErr(err)
	//	}
	//}
}
