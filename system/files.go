package system

import (
	"os"
	"os/user"
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
	currentUser, err := user.Current()
	if err != nil {
		return currentUser.HomeDir
	}else{
		return ""
	}
}