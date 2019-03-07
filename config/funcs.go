package config

import (
	"encoding/json"
	"errors"
	"github.com/jfrog/jfrog-client-go/utils/errorutils"
	"github.com/jfrog/jfrog-client-go/utils/io/fileutils"
	"os"
	"path/filepath"
)

func GetEnvxHomeDir() (string, error) {

	if os.Getenv(EnvxHomeDirEnv) != "" {
		return os.Getenv(EnvxHomeDirEnv), nil
	}

	userHomeDir := fileutils.GetHomeDir()
	if userHomeDir == "" {
		err := errorutils.CheckError(errors.New("couldn't find home directory, make sure your HOME environment variable is set"))
		if err != nil {
			return "", err
		}
	}
	return filepath.Join(userHomeDir, EnvxDirNameDefault), nil
}

func ReadConfig() (*EnvxConfigV0, error) {
	confFilePath, err := getConfFilePath()
	if err != nil {
		return nil, err
	}
	conf := new(EnvxConfigV0)
	exists, err := fileutils.IsFileExists(confFilePath, false)
	if err != nil {
		return nil, err
	}
	if !exists {
		return conf, nil
	}
	content, err := fileutils.ReadFile(confFilePath)
	if err != nil {
		return nil, err
	}
	if len(content) == 0 {
		return new(EnvxConfigV0), nil
	}
	err = json.Unmarshal(content, &conf)
	return conf, err
}

func getConfFilePath() (string, error) {
	confPath, err := GetEnvxHomeDir()
	if err != nil {
		return "", err
	}
	err = os.MkdirAll(confPath, 0777)
	if err != nil {
		return "", err
	}
	return filepath.Join(confPath, EnvxConfigFile), nil
}
