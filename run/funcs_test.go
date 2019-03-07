package run

import (
	"runtime"
	"testing"
)

func TestRun(t *testing.T) {
	var command string
	var args []string
	var envVars []string
	var printVars []string
	if runtime.GOOS == "windows" {
		command = "cmd"
		args = []string{"/c", "echo", "TESTVAR=%TESTVAR%"}
	} else {
		command = "/bin/sh"
		args = []string{"-c", "echo", "TESTVAR=$TESTVAR"}
	}

	_, _, err := Run(command, args, envVars, printVars)
	if err != nil {
		t.Fatal()
	}

	//t.Logf("result %s", result)
	//if !strings.HasPrefix(result, expected) {
	//	t.Errorf("result (%s) does not match expected (%s)", result, expected)
	//}

}

func TestRunWithEnv(t *testing.T) {
	var command string
	var args []string
	envVars := []string{"TESTVAR=TESTVAL"}
	var printVars []string
	if runtime.GOOS == "windows" {
		command = "cmd"
		args = []string{"/c", "echo", "TESTVAR=%TESTVAR%"}
	} else {
		command = "/bin/sh"
		args = []string{"-c", "echo", "TESTVAR=$TESTVAR"}
	}

	_, _, err := Run(command, args, envVars, printVars)
	if err != nil {
		t.Fatal()
	}

	//t.Logf("result %s", result)
	//if !strings.HasPrefix(result, expected) {
	//	t.Errorf("result (%s) does not match expected (%s)", result, expected)
	//}

}
func TestRunPrintVarsNoEnv(t *testing.T) {
	var command string
	var args []string
	var envVars []string
	printVars := []string{"TESTVAR", "PATH"}
	if runtime.GOOS == "windows" {
		command = "cmd"
		args = []string{"/c", "echo", "TESTVAR=%TESTVAR%"}
	} else {
		command = "/bin/sh"
		args = []string{"-c", "echo", "TESTVAR=$TESTVAR"}
	}

	_, _, err := Run(command, args, envVars, printVars)
	if err != nil {
		t.Fatal()
	}

	//t.Logf("result %s", result)
	//if !strings.HasPrefix(result, expected) {
	//	t.Errorf("result (%s) does not match expected (%s)", result, expected)
	//}

}

func TestRunPrintVarsWithEnv(t *testing.T) {
	var command string
	var args []string
	envVars := []string{"TESTVAR=TESTVAL"}
	printVars := []string{"TESTVAR", "PATH"}
	if runtime.GOOS == "windows" {
		command = "cmd"
		args = []string{"/c", "echo", "TESTVAR=%TESTVAR%"}
	} else {
		command = "/bin/sh"
		args = []string{"-c", "echo", "TESTVAR=$TESTVAR"}
	}

	_, _, err := Run(command, args, envVars, printVars)
	if err != nil {
		t.Fatal()
	}

	//t.Logf("result %s", result)
	//if !strings.HasPrefix(result, expected) {
	//	t.Errorf("result (%s) does not match expected (%s)", result, expected)
	//}

}
