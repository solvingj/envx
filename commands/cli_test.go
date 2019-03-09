package commands

import (
	"bytes"
	"flag"
	"github.com/codegangsta/cli"
	"os"
	"runtime"
	"strings"
	"testing"
)

func runAppSubCmd(command string) (string, error) {
	args := strings.Split(command, " ")
	app := GetApp()
	outBuffer := new(bytes.Buffer)
	app.Writer = outBuffer
	parFlagSet := flag.NewFlagSet("", 0)
	parContext := cli.NewContext(app, parFlagSet, nil)

	flagSet := flag.NewFlagSet("", 0)
	flagSet.Parse(args)
	context := cli.NewContext(app, flagSet, parContext)
	err := app.RunAsSubcommand(context)

	if err != nil {
		return "", err
	}
	outBytes := outBuffer.Bytes()
	result := string(outBytes[:])
	return result, nil
}

func TestCommandNoArgs(t *testing.T) {
	command := "envx"
	result, err := runAppSubCmd(command)
	if err != nil {
		t.Fatal()
	}

	t.Log(result)

	//expected := fmt.Sprintf(`:
   //%s `, command)
   //
	//if !strings.Contains(result, expected) {
	//	t.Errorf("result (%s) does not contain expected (%s)", result, expected)
	//}
}

func TestCommandConfig(t *testing.T) {
	command := "envx config"
	result, err := runAppSubCmd(command)
	if err != nil {
		t.Fatal()
	}

	t.Log(result)

	//expected := fmt.Sprintf(`USAGE:
   //%s `, command)
   //
	//if !strings.Contains(result, expected) {
	//	t.Errorf("result (%s) does not contain expected (%s)", result, expected)
	//}
}

func TestCommandEnv(t *testing.T) {
	command := "envx env"
	result, err := runAppSubCmd(command)
	if err != nil {
		t.Fatal()
	}

	t.Log(result)

	//expected := fmt.Sprintf(`USAGE:
   //%s `, command)
	//
	//if !strings.Contains(result, expected) {
	//	t.Errorf("result (%s) does not contain expected (%s)", result, expected)
	//}
}

func TestCommandEnvList(t *testing.T) {
	command := "envx env list"
	result, err := runAppSubCmd(command)
	if err != nil {
		t.Fatal()
	}

	t.Log(result)

	//expected := fmt.Sprintf(`USAGE:
   //%s `, command)
   //
	//if !strings.Contains(result, expected) {
	//	t.Errorf("result (%s) does not contain expected (%s)", result, expected)
	//}
}

func TestCommandRunEcho(t *testing.T) {
	var commandArg string
	if runtime.GOOS == "windows" {
		commandArg = `cmd /c echo TESTVAR=%TESTVAR%`
	} else {
		commandArg = `/bin/sh -c echo TESTVAR=$TESTVAR`
	}

	command := "envx run " + commandArg
	t.Log(command)
	result, err := runAppSubCmd(command)
	if err != nil {
		print(err.Error())
		t.Fatal()
	}

	t.Log(result)

	//expected := fmt.Sprintf(`USAGE:
   //%s `, command)
   //
	//if !strings.Contains(result, expected) {
	//	t.Errorf("result (%s) does not contain expected (%s)", result, expected)
	//}
}

func TestCommandEnvShow(t *testing.T) {
	_ = os.Setenv("ENVX_HOME_DIR", "test_data")
	command := "envx env show test_env_01"
	result, err := runAppSubCmd(command)

	if err != nil {
		t.Fatal()
	}

	t.Log(result)

	//expected := fmt.Sprintf(`USAGE:
   //%s `, command)
   //
	//if !strings.Contains(result, expected) {
	//	t.Errorf("result (%s) does not contain expected (%s)", result, expected)
	//}
}

