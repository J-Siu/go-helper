package helper

import (
	"os"
	"path"
)

// Get full path of current directory.
// Return string pointer.
func CurrentPath() *string {
	p, _ := os.Getwd()
	return &p
}

// Get current directory name. Not full path.
// Return string pointer.
func CurrentDirBase() *string {
	p := path.Base(*CurrentPath())
	return &p
}

// Return full path of path provided.
// If <workpathP> is empty/nil, use current path.
// <workpathP> not modified.
// Return string pointer.
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
