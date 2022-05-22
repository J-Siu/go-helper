package helper

import (
	"os"
	"path"
	"strings"
)

// Get full path of current directory.
//  - Return string pointer.
func CurrentPath() *string {
	p, _ := os.Getwd()
	return &p
}

// Get current directory name. Not full path.
//  - Return string pointer.
func CurrentDirBase() *string {
	p := path.Base(*CurrentPath())
	return &p
}

// Return full path of path provided.
//  - If <workPathP> is empty/nil, use current path.
//  - <workPathP> not modified.
//  - Return string pointer.
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

// Return full path of path provided.
//  - If <workPath> is empty, use current path.
//  - Return string pointer.
func FullPathStr(workPath string) *string {
	return FullPath(&workPath)
}

// Check workPath is regular file
func IsRegularFile(workPath string) bool {
	fileInfo, err := os.Stat(workPath)
	if err != nil {
		if Debug {
			ReportDebug(fileInfo, "IsRegularFile:fileInfo", false, false)
			Errs = append(Errs, err)
		}
		return false
	}
	return fileInfo.Mode().IsRegular()
}

// Check workPath is directory
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

// Check two paths have same parent directory
func SameDir(path1, path2 string) bool {
	return path.Dir(path1) == path.Dir(path2)
}

// Check file has supplied extension
func FileHasExt(name, ext string) bool {
	return strings.ToLower(path.Ext(name)) == strings.ToLower(ext)
}

// Return filename/path with extension removed
func FileRemoveExt(filename string) string {
	return strings.TrimSuffix(filename, path.Ext(filename))
}

// Search for file in a directory
//  - filename path info removed automatically, only base is used
//  - case insensitive search
//  - return actually filename if found
//  - return empty if not found or error
//  - error is added to Errs
func FileInDir(dir, filename string) string {
	var fileBase string = strings.ToLower(path.Base(filename))
	dirEntry, err := os.ReadDir(dir)
	if err != nil {
		Errs = append(Errs, err)
		return ""
	}
	for _, f := range dirEntry {
		if fileBase == strings.ToLower(f.Name()) {
			// case insensitive matched, return real name
			return f.Name()
		}
	}
	return ""
}

// Apply following to filename:
//	- remove extension
//	- to lowercase
//	- remove -,_
func FileSimplifyName(filename string) string {
	filename = FileRemoveExt(filename)
	filename = strings.ToLower(filename)
	filename = strings.ReplaceAll(filename, "-", "")
	filename = strings.ReplaceAll(filename, "_", "")
	return filename
}
