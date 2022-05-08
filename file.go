package helper

import (
	"os"
	"path"
)

// Get full path of current directory.
func CurrentPath() string {
	fullpath, _ := os.Getwd()
	return fullpath
}

// Get current directory name. Not full path.
func CurrentDirBase() string {
	return path.Base(CurrentPath())
}
