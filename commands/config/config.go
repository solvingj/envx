package config

import (
	"errors"
	"github.com/codegangsta/cli"
	"github.com/solvingj/envx/commands/common"
	configdocs "github.com/solvingj/envx/docs/envx/config"
	"github.com/solvingj/envx/system"
)

func GetCommands() []cli.Command {
	return []cli.Command{
		{
			Name:      "show",
			Flags:     getFlags(),
			Usage:     configdocs.DescriptionShow,
			HideHelp:  true,
			Action: func(c *cli.Context) {
				configCmd(c)
			},
		},
	}
}

func getFlags() []cli.Flag {
	var flags []cli.Flag
	return flags
}

//TODO: Implement configCmd function
func configCmd(c *cli.Context) {
	err := errors.New("command not implemented")
	if err != nil {
		system.ExitOnErr(err)
	}
	if len(c.Args()) < 2 {
		common.PrintHelpAndExitWithError("Wrong number of arguments.", c)
	}

	//configCommandConfiguration := commands.CreateCommandConfig(c)

	//if len(c.Args()) > 0 {
	//	if c.Args()[2] == "list" {
	//		commands.ListConfigs(configCommandConfiguration)
	//		return
	//	} else if c.Args()[0] == "show" {
	//		commands.ShowConfig(configCommandConfiguration.Interactive)
	//		return
	//
	//		system.ExitOnErr(err)
	//	}
	//}
}
