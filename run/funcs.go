package run

import (
	"bytes"
	"fmt"
	"github.com/jfrog/jfrog-client-go/utils/log"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func Run(command string, args []string, envVars []string, printVars[]string)(outStr string, errStr string, err error){
	cmd := exec.Command(command, args...)
	cmd.Env = envVars

	if  len(printVars) > 0 {
		log.Output("----Begin : Logged Environment Variables----")
	}

	for _, printEnvVar := range printVars{
		log.Debug("Looking for variable:" + printEnvVar)
		for _, useEnvVar := range envVars{
			useEnvVarSplit := strings.Split(useEnvVar, "=")
			k := useEnvVarSplit[0]
			v := useEnvVarSplit[1]
			if printEnvVar == k {
				log.Output(k + "=" + v)
				log.Debug("Found var in passed-in-list of env vars: " + k)
				break
			}else{
				envVar := os.Getenv(printEnvVar)
				if envVar != ""{
					log.Output(printEnvVar + "=" + envVar)
					log.Debug("Found env var environment: " + envVar)
					break
				}
			}
			log.Debug("Did not find requested environment variable: " + printEnvVar)
		}
	}
	if len(printVars) > 0 {
		log.Output("----End   : Logged Environment Variables----")
	}

	var stdoutBuf, stderrBuf bytes.Buffer
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err = cmd.Start()
	if err != nil {
		log.Error(fmt.Sprintf("cmd.Start() failed with '%s'\n", err))
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
		wg.Done()
	}()

	_, errStderr = io.Copy(stderr, stderrIn)
	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		log.Error(fmt.Sprintf("cmd.Run() failed with %s\n", err))
	}
	if errStdout != nil || errStderr != nil {
		log.Error("failed to capture stdout or stderr\n")
	}
	outStr, errStr = string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	log.Debug(fmt.Sprintf("\nout:\n%s\nerr:\n%s\n", outStr, errStr))
	return outStr, errStr, err
}
