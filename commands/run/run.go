package run

import (
	"github.com/codegangsta/cli"
	"github.com/solvingj/envx/commands/common"
	"github.com/solvingj/envx/env"
	"github.com/solvingj/envx/run"
)

func RunCmd(c *cli.Context) {

	if len(c.Args()) < 1 {
		common.PrintHelpAndExitWithError("Wrong number of arguments.", c)
	}
	cmd := c.Args().First()
	cmdArgs := c.Args().Tail()
	envArgs := c.StringSlice("with-env")
	envVarArgs := c.StringSlice("e")
	printVars := c.StringSlice("print-vars")

	var envVarsToUse []string
	for _, envName := range envArgs {
		environ, err := env.ReadEnv(envName)
		if err != nil {
			common.ExitOnErr(err)
		}
		for k, v := range environ.Vars {
			envVarsToUse = append(envVarsToUse, k+"="+v)
			if err != nil {
				common.ExitOnErr(err)
			}
		}
	}
	for _, envVar := range envVarArgs {
		envVarsToUse = append(envVarsToUse, envVar)
	}

	_, _, err := run.Run(cmd, cmdArgs, envVarsToUse, printVars)
	if err != nil {
		common.ExitOnErr(err)
	}
}

func Flags() []cli.Flag {
	var flags []cli.Flag
	flags = append(flags, EnvVarFlag()...)
	flags = append(flags, WithEnvFlag()...)
	return flags
}

func EnvVarFlag() []cli.Flag {
	return []cli.Flag{
		cli.StringSliceFlag{
			Name:  "e",
			Usage: "List of environment variables in the form of \"key1=value1\"` `",
		},
	}
}

func WithEnvFlag() []cli.Flag {
	return []cli.Flag{
		cli.StringSliceFlag{
			Name:  "with-env",
			Usage: "List of environments to apply prior to running command` `",
		},
	}
}
