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
// If <workPathP> is empty/nil, use current path.
// <workPathP> not modified.
// Return string pointer.
func FullPath(workPathP *string) *string {
	var p string
	if workPathP == nil || *workPathP == "" {
		// Empty return current path
		p = *CurrentPath()
	} else if (*workPathP)[0] == '/' {
		// Path start with / already full path, return it
		p = *workPathP
	} else {
		// Add to current path
		p = path.Join(*CurrentPath(), *workPathP)
	}
	return &p
}

func IsRegularFile(workPath string) bool {
	fileInfo, err := os.Stat(workPath)
	if err != nil {
		if Debug {
			Errs = append(Errs, err)
		}
		return false
	}
	return fileInfo.Mode().IsRegular()
}

func IsDir(workPath string) bool {
	fileInfo, err := os.Stat(workPath)
	if err != nil {
		if Debug {
			Errs = append(Errs, err)
		}
		return false
	}
	return fileInfo.IsDir()
}

func SameDir(path1, path2 string) bool {
	return path.Dir(path1) == path.Dir(path2)
}
