package system

import (
	"github.com/mitchellh/go-homedir"
	"os"
)
// FileExists reports whether the named file exists as a boolean
func FileExists(name string) bool {
	if fi, err := os.Stat(name); err == nil {
		if fi.Mode().IsRegular() {
			return true
		}
	}
	return false
}

func DirExists(name string) bool {
	if fi, err := os.Stat(name); err == nil {
		if fi.Mode().IsDir() {
			return true
		}
	}
	return false
}

func HomeDir() string {
	homeDir, err := homedir.Dir()
	if err != nil {
		return homeDir
	}else{
		return ""
	}
}
