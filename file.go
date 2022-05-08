package helper

import (
	"os"
	"path"
)

// Get full path of current directory.
func CurrentPath() *string {
	p, _ := os.Getwd()
	return &p
}

// Get current directory name. Not full path.
func CurrentDirBase() *string {
	p := path.Base(*CurrentPath())
	return &p
}

// Return full path of path provided.
// An Empty 'workpath' will return current path.
// 'workpath' not modified.
func FullPath(workpathP *string) *string {
	var p string
	if workpathP == nil || *workpathP == "" {
		// Empty return current path
		p = *CurrentPath()
	} else if (*workpathP)[0] == '/' {
		// Path start with / already full path, return it
		p = *workpathP
	} else {
		// Add to current path
		p = path.Join(*CurrentPath(), *workpathP)
	}
	return &p
}
