package env

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/solvingj/envx/config"
	"github.com/solvingj/envx/system"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func getEnvxEnvDir() (string, error) {
	envxHomeDir, err := config.GetEnvxHomeDir()
	if err != nil {
		return "", err
	} else {
		return filepath.Join(envxHomeDir, EnvxEnvDirDefault), err
	}
}

func ReadEnv(envName string) (*Env, error) {
	envDir, err := getEnvxEnvDir()
	if err != nil {
		return nil, err
	}
	var envPath = filepath.Join(envDir, envName)

	var envPathFull = ""
	for _, ext := range []string{"env", "envx", "json", "yaml"} {
		envPathFullCurrent := envPath + "." + ext
		exists := system.FileExists(envPathFullCurrent)
		if exists {
			envPathFull = envPathFullCurrent
			break
		}
	}

	if envPathFull == "" {
		return nil, errors.New(fmt.Sprintf("No environment found with name: %s in directory %s", envName, envDir))
	}
	content, err := ioutil.ReadFile(envPathFull)
	if err != nil {
		return nil, err
	}
	if len(content) == 0 {
		return nil, errors.New(fmt.Sprintf("Environment found but contents were empty: %s", envPathFull))
	}
	envMap, err := godotenv.Unmarshal(string(content))
	if err != nil {
		return nil, err
	}
	env := new(Env)
	env.Vars = envMap
	env.Path = envPathFull
	env.IsDefault = false
	return env, err
}

func EnumerateEnvs() ([]string, error) {
	envxHomeDir, err := config.GetEnvxHomeDir()
	var envFullPath = filepath.Join(envxHomeDir, "env")
	log.Println(fmt.Sprintf("envFullPath = %s", envFullPath))

	files, err := ioutil.ReadDir(envFullPath)

	if err != nil {
		log.Fatal(err)
	}
	var envNames []string

	for _, file := range files {
		log.Println(fmt.Sprintf("file = %s", file.Name()))
		filePrefix := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		log.Println(fmt.Sprintf("filePrefix = %s", filePrefix))
		fileExt := filepath.Ext(file.Name())
		log.Println(fmt.Sprintf("fileExt = %s", fileExt))
		if fileExt == ".envx" {
			envNames = append(envNames, filePrefix)
		}
		if fileExt == ".env" {
			envNames = append(envNames, filePrefix)
		}
		if fileExt == ".yaml" {
			envNames = append(envNames, filePrefix)
		}
	}
	log.Println(fmt.Sprintf("envNames = %s", envNames))

	return envNames, nil
}

func ShowEnv(envId string) error {
	err := printEnv(envId)
	return err
}

func ListEnvs() error {
	envNames, err := EnumerateEnvs()
	if err != nil {
		return err
	} else {
		for _, env := range envNames {
			fmt.Println(env)
		}
		return nil
	}
}

func printEnv(envId string) error {
	envContents, err := ReadEnv(envId)
	if err != nil {
		return err
	} else {
		for key, val := range envContents.Vars {
			log.Println(key + "=" + val)
		}
		return nil
	}
}
